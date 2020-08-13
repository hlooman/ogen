// Code generated by fastssz. DO NOT EDIT.
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Block object
func (b *Block) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the Block object to a target array
func (b *Block) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(596)

	// Field (0) 'Header'
	if b.Header == nil {
		b.Header = new(BlockHeader)
	}
	if dst, err = b.Header.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'Votes'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Votes); ii++ {
		offset += 4
		offset += b.Votes[ii].SizeSSZ()
	}

	// Offset (2) 'Txs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Txs) * 188

	// Offset (3) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Deposits) * 308

	// Offset (4) 'Exits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Exits) * 192

	// Offset (5) 'VoteSlashings'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.VoteSlashings); ii++ {
		offset += 4
		offset += b.VoteSlashings[ii].SizeSSZ()
	}

	// Offset (6) 'RANDAOSlashings'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.RANDAOSlashings) * 152

	// Offset (7) 'ProposerSlashings'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ProposerSlashings) * 984

	// Offset (8) 'GovernanceVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.GovernanceVotes) * 260

	// Field (9) 'Signature'
	dst = append(dst, b.Signature[:]...)

	// Field (10) 'RandaoSignature'
	dst = append(dst, b.RandaoSignature[:]...)

	// Field (1) 'Votes'
	if len(b.Votes) > 32 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.Votes)
		for ii := 0; ii < len(b.Votes); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Votes[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Votes); ii++ {
		if dst, err = b.Votes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (2) 'Txs'
	if len(b.Txs) > 5000 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Txs); ii++ {
		if dst, err = b.Txs[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'Deposits'
	if len(b.Deposits) > 128 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'Exits'
	if len(b.Exits) > 128 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.Exits); ii++ {
		if dst, err = b.Exits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (5) 'VoteSlashings'
	if len(b.VoteSlashings) > 10 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.VoteSlashings)
		for ii := 0; ii < len(b.VoteSlashings); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.VoteSlashings[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.VoteSlashings); ii++ {
		if dst, err = b.VoteSlashings[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (6) 'RANDAOSlashings'
	if len(b.RANDAOSlashings) > 20 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.RANDAOSlashings); ii++ {
		if dst, err = b.RANDAOSlashings[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (7) 'ProposerSlashings'
	if len(b.ProposerSlashings) > 2 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.ProposerSlashings); ii++ {
		if dst, err = b.ProposerSlashings[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (8) 'GovernanceVotes'
	if len(b.GovernanceVotes) > 128 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.GovernanceVotes); ii++ {
		if dst, err = b.GovernanceVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Block object
func (b *Block) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 596 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2, o3, o4, o5, o6, o7, o8 uint64

	// Field (0) 'Header'
	if b.Header == nil {
		b.Header = new(BlockHeader)
	}
	if err = b.Header.UnmarshalSSZ(buf[0:372]); err != nil {
		return err
	}

	// Offset (1) 'Votes'
	if o1 = ssz.ReadOffset(buf[372:376]); o1 > size {
		return ssz.ErrOffset
	}

	// Offset (2) 'Txs'
	if o2 = ssz.ReadOffset(buf[376:380]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Offset (3) 'Deposits'
	if o3 = ssz.ReadOffset(buf[380:384]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'Exits'
	if o4 = ssz.ReadOffset(buf[384:388]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Offset (5) 'VoteSlashings'
	if o5 = ssz.ReadOffset(buf[388:392]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Offset (6) 'RANDAOSlashings'
	if o6 = ssz.ReadOffset(buf[392:396]); o6 > size || o5 > o6 {
		return ssz.ErrOffset
	}

	// Offset (7) 'ProposerSlashings'
	if o7 = ssz.ReadOffset(buf[396:400]); o7 > size || o6 > o7 {
		return ssz.ErrOffset
	}

	// Offset (8) 'GovernanceVotes'
	if o8 = ssz.ReadOffset(buf[400:404]); o8 > size || o7 > o8 {
		return ssz.ErrOffset
	}

	// Field (9) 'Signature'
	copy(b.Signature[:], buf[404:500])

	// Field (10) 'RandaoSignature'
	copy(b.RandaoSignature[:], buf[500:596])

	// Field (1) 'Votes'
	{
		buf = tail[o1:o2]
		num, err := ssz.DecodeDynamicLength(buf, 32)
		if err != nil {
			return err
		}
		b.Votes = make([]*MultiValidatorVote, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Votes[indx] == nil {
				b.Votes[indx] = new(MultiValidatorVote)
			}
			if err = b.Votes[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (2) 'Txs'
	{
		buf = tail[o2:o3]
		num, err := ssz.DivideInt2(len(buf), 188, 5000)
		if err != nil {
			return err
		}
		b.Txs = make([]*Tx, num)
		for ii := 0; ii < num; ii++ {
			if b.Txs[ii] == nil {
				b.Txs[ii] = new(Tx)
			}
			if err = b.Txs[ii].UnmarshalSSZ(buf[ii*188 : (ii+1)*188]); err != nil {
				return err
			}
		}
	}

	// Field (3) 'Deposits'
	{
		buf = tail[o3:o4]
		num, err := ssz.DivideInt2(len(buf), 308, 128)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		for ii := 0; ii < num; ii++ {
			if b.Deposits[ii] == nil {
				b.Deposits[ii] = new(Deposit)
			}
			if err = b.Deposits[ii].UnmarshalSSZ(buf[ii*308 : (ii+1)*308]); err != nil {
				return err
			}
		}
	}

	// Field (4) 'Exits'
	{
		buf = tail[o4:o5]
		num, err := ssz.DivideInt2(len(buf), 192, 128)
		if err != nil {
			return err
		}
		b.Exits = make([]*Exit, num)
		for ii := 0; ii < num; ii++ {
			if b.Exits[ii] == nil {
				b.Exits[ii] = new(Exit)
			}
			if err = b.Exits[ii].UnmarshalSSZ(buf[ii*192 : (ii+1)*192]); err != nil {
				return err
			}
		}
	}

	// Field (5) 'VoteSlashings'
	{
		buf = tail[o5:o6]
		num, err := ssz.DecodeDynamicLength(buf, 10)
		if err != nil {
			return err
		}
		b.VoteSlashings = make([]*VoteSlashing, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.VoteSlashings[indx] == nil {
				b.VoteSlashings[indx] = new(VoteSlashing)
			}
			if err = b.VoteSlashings[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (6) 'RANDAOSlashings'
	{
		buf = tail[o6:o7]
		num, err := ssz.DivideInt2(len(buf), 152, 20)
		if err != nil {
			return err
		}
		b.RANDAOSlashings = make([]*RANDAOSlashing, num)
		for ii := 0; ii < num; ii++ {
			if b.RANDAOSlashings[ii] == nil {
				b.RANDAOSlashings[ii] = new(RANDAOSlashing)
			}
			if err = b.RANDAOSlashings[ii].UnmarshalSSZ(buf[ii*152 : (ii+1)*152]); err != nil {
				return err
			}
		}
	}

	// Field (7) 'ProposerSlashings'
	{
		buf = tail[o7:o8]
		num, err := ssz.DivideInt2(len(buf), 984, 2)
		if err != nil {
			return err
		}
		b.ProposerSlashings = make([]*ProposerSlashing, num)
		for ii := 0; ii < num; ii++ {
			if b.ProposerSlashings[ii] == nil {
				b.ProposerSlashings[ii] = new(ProposerSlashing)
			}
			if err = b.ProposerSlashings[ii].UnmarshalSSZ(buf[ii*984 : (ii+1)*984]); err != nil {
				return err
			}
		}
	}

	// Field (8) 'GovernanceVotes'
	{
		buf = tail[o8:]
		num, err := ssz.DivideInt2(len(buf), 260, 128)
		if err != nil {
			return err
		}
		b.GovernanceVotes = make([]*GovernanceVote, num)
		for ii := 0; ii < num; ii++ {
			if b.GovernanceVotes[ii] == nil {
				b.GovernanceVotes[ii] = new(GovernanceVote)
			}
			if err = b.GovernanceVotes[ii].UnmarshalSSZ(buf[ii*260 : (ii+1)*260]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Block object
func (b *Block) SizeSSZ() (size int) {
	size = 596

	// Field (1) 'Votes'
	for ii := 0; ii < len(b.Votes); ii++ {
		size += 4
		size += b.Votes[ii].SizeSSZ()
	}

	// Field (2) 'Txs'
	size += len(b.Txs) * 188

	// Field (3) 'Deposits'
	size += len(b.Deposits) * 308

	// Field (4) 'Exits'
	size += len(b.Exits) * 192

	// Field (5) 'VoteSlashings'
	for ii := 0; ii < len(b.VoteSlashings); ii++ {
		size += 4
		size += b.VoteSlashings[ii].SizeSSZ()
	}

	// Field (6) 'RANDAOSlashings'
	size += len(b.RANDAOSlashings) * 152

	// Field (7) 'ProposerSlashings'
	size += len(b.ProposerSlashings) * 984

	// Field (8) 'GovernanceVotes'
	size += len(b.GovernanceVotes) * 260

	return
}

// HashTreeRoot ssz hashes the Block object
func (b *Block) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the Block object with a hasher
func (b *Block) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Header'
	if err = b.Header.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Votes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Votes))
		if num > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Votes[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 32)
	}

	// Field (2) 'Txs'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Txs))
		if num > 5000 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Txs[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 5000)
	}

	// Field (3) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Deposits[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 128)
	}

	// Field (4) 'Exits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Exits))
		if num > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.Exits[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 128)
	}

	// Field (5) 'VoteSlashings'
	{
		subIndx := hh.Index()
		num := uint64(len(b.VoteSlashings))
		if num > 10 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.VoteSlashings[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 10)
	}

	// Field (6) 'RANDAOSlashings'
	{
		subIndx := hh.Index()
		num := uint64(len(b.RANDAOSlashings))
		if num > 20 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.RANDAOSlashings[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 20)
	}

	// Field (7) 'ProposerSlashings'
	{
		subIndx := hh.Index()
		num := uint64(len(b.ProposerSlashings))
		if num > 2 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.ProposerSlashings[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2)
	}

	// Field (8) 'GovernanceVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.GovernanceVotes))
		if num > 128 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = b.GovernanceVotes[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 128)
	}

	// Field (9) 'Signature'
	hh.PutBytes(b.Signature[:])

	// Field (10) 'RandaoSignature'
	hh.PutBytes(b.RandaoSignature[:])

	hh.Merkleize(indx)
	return
}