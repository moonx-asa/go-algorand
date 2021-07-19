// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package merklekeystore

import (
	"fmt"
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/crypto/merklearray"
)

type (
	//EphemeralKeys represent the possible keys inside the keystore.
	//msgp:allocbound EphemeralKeys -
	EphemeralKeys []crypto.SignatureAlgorithm

	//Proof represent the merkle proof in each signature.
	//msgp:allocbound Proof -
	Proof []crypto.Digest
)

//Length returns the amount of disposable keys
func (d EphemeralKeys) Length() uint64 {
	return uint64(len(d))
}

// GetHash Gets the hash of the VerifyingKey tied to the signatureAlgorithm in pos.
func (d EphemeralKeys) GetHash(pos uint64) (crypto.Digest, error) {
	return disposableKeyHash(&d[pos])
}

func disposableKeyHash(s *crypto.SignatureAlgorithm) (crypto.Digest, error) {
	vkey := s.GetSigner().GetVerifyingKey()
	return crypto.HashObj(&vkey), nil
}

// Signature is a byte signature on a crypto.Hashable object, and includes a merkle proof for the signing key.
type Signature struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	crypto.ByteSignature `codec:"bsig"`
	Proof                `codec:"prf"`
	VerifyingKey         crypto.VerifyingKey `codec:"vkey"`
	Pos                  uint64              `codec:"pos"`
}

// Signer is a merkleKeyStore, contain multiple keys which can be used per round.
type Signer struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`
	// these keys are the keys used to sign in a round.
	// should be disposed of once possible.
	EphemeralKeys    `codec:"keys,allocbound=-"`
	StartRound       uint64 `codec:"sround"`
	merklearray.Tree `codec:"tree"`
}

var errStartBiggerThanEndRound = fmt.Errorf("cannot create merkleKeyStore because end round is smaller then start round")

// New Generates a merklekeystore.Signer
// Note that the signer will have keys for the rounds  [StartRound, endRound)
func New(startRound, endRound uint64, sigAlgoType crypto.AlgorithmType) (*Signer, error) {
	if startRound >= endRound {
		return nil, errStartBiggerThanEndRound
	}
	keys := make(EphemeralKeys, endRound-startRound)
	for i := range keys {
		keys[i] = *crypto.NewSigner(sigAlgoType)
	}
	tree, err := merklearray.Build(keys)
	if err != nil {
		return nil, err
	}

	return &Signer{
		EphemeralKeys: keys,
		StartRound:    startRound,
		Tree:          *tree,
	}, nil
}

// GetVerifier can be used to store the commitment and verifier for this signer.
func (m *Signer) GetVerifier() *Verifier {
	return &Verifier{
		root: m.Root(),
	}
}

// Sign outputs a signature + proof for the signing key.
func (m *Signer) Sign(hashable crypto.Hashable, round int) (Signature, error) {
	pos, err := m.getKeyPosition(uint64(round))
	if err != nil {
		return Signature{}, err
	}

	proof, err := m.Prove([]uint64{pos})
	if err != nil {
		return Signature{}, err
	}

	signer := m.EphemeralKeys[pos].GetSigner()
	return Signature{
		ByteSignature: signer.Sign(hashable),
		Proof:         proof,
		VerifyingKey:  signer.GetVerifyingKey(),
		Pos:           pos,
	}, nil
}

var errOutOfBounds = fmt.Errorf("cannot find signing key for given round")

func (m *Signer) getKeyPosition(round uint64) (uint64, error) {
	if round < m.StartRound {
		return 0, errOutOfBounds
	}

	pos := round - m.StartRound
	if pos >= uint64(len(m.EphemeralKeys)) {
		return 0, errOutOfBounds
	}
	return pos, nil
}