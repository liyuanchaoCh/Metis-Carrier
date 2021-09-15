package taskmng

import (
	"github.com/RosettaFlow/Carrier-Go/lib/netmsg/common"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the TaskResultMsg object
func (t *TaskResultMsg) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(t)
}

// MarshalSSZTo ssz marshals the TaskResultMsg object to a target array
func (t *TaskResultMsg) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20)

	// Offset (0) 'MsgOption'
	dst = ssz.WriteOffset(dst, offset)
	if t.MsgOption == nil {
		t.MsgOption = new(common.MsgOption)
	}
	offset += t.MsgOption.SizeSSZ()

	// Offset (1) 'TaskEventList'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(t.TaskEventList); ii++ {
		offset += 4
		offset += t.TaskEventList[ii].SizeSSZ()
	}

	// Field (2) 'GetCreateAt'
	dst = ssz.MarshalUint64(dst, t.CreateAt)

	// Offset (3) 'Sign'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.Sign)

	// Field (0) 'MsgOption'
	if dst, err = t.MsgOption.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'TaskEventList'
	if len(t.TaskEventList) > 16777216 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(t.TaskEventList)
		for ii := 0; ii < len(t.TaskEventList); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += t.TaskEventList[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(t.TaskEventList); ii++ {
		if dst, err = t.TaskEventList[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'Sign'
	if len(t.Sign) > 1024 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, t.Sign...)

	return
}

// UnmarshalSSZ ssz unmarshals the TaskResultMsg object
func (t *TaskResultMsg) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o3 uint64

	// Offset (0) 'MsgOption'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'TaskEventList'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'GetCreateAt'
	t.CreateAt = ssz.UnmarshallUint64(buf[8:16])

	// Offset (3) 'Sign'
	if o3 = ssz.ReadOffset(buf[16:20]); o3 > size || o1 > o3 {
		return ssz.ErrOffset
	}

	// Field (0) 'MsgOption'
	{
		buf = tail[o0:o1]
		if t.MsgOption == nil {
			t.MsgOption = new(common.MsgOption)
		}
		if err = t.MsgOption.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'TaskEventList'
	{
		buf = tail[o1:o3]
		num, err := ssz.DecodeDynamicLength(buf, 16777216)
		if err != nil {
			return err
		}
		t.TaskEventList = make([]*TaskEvent, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if t.TaskEventList[indx] == nil {
				t.TaskEventList[indx] = new(TaskEvent)
			}
			if err = t.TaskEventList[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'Sign'
	{
		buf = tail[o3:]
		if len(buf) > 1024 {
			return ssz.ErrBytesLength
		}
		if cap(t.Sign) == 0 {
			t.Sign = make([]byte, 0, len(buf))
		}
		t.Sign = append(t.Sign, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the TaskResultMsg object
func (t *TaskResultMsg) SizeSSZ() (size int) {
	size = 20

	// Field (0) 'MsgOption'
	if t.MsgOption == nil {
		t.MsgOption = new(common.MsgOption)
	}
	size += t.MsgOption.SizeSSZ()

	// Field (1) 'TaskEventList'
	for ii := 0; ii < len(t.TaskEventList); ii++ {
		size += 4
		size += t.TaskEventList[ii].SizeSSZ()
	}

	// Field (3) 'Sign'
	size += len(t.Sign)

	return
}

// HashTreeRoot ssz hashes the TaskResultMsg object
func (t *TaskResultMsg) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(t)
}

// HashTreeRootWith ssz hashes the TaskResultMsg object with a hasher
func (t *TaskResultMsg) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'MsgOption'
	if err = t.MsgOption.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'TaskEventList'
	{
		subIndx := hh.Index()
		num := uint64(len(t.TaskEventList))
		if num > 16777216 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = t.TaskEventList[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16777216)
	}

	// Field (2) 'GetCreateAt'
	hh.PutUint64(t.CreateAt)

	// Field (3) 'Sign'
	if len(t.Sign) > 1024 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(t.Sign)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the TaskEvent object
func (t *TaskEvent) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(t)
}

// MarshalSSZTo ssz marshals the TaskEvent object to a target array
func (t *TaskEvent) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(24)

	// Offset (0) 'Type'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.Type)

	// Offset (1) 'TaskId'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.TaskId)

	// Offset (2) 'IdentityId'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.IdentityId)

	// Offset (3) 'Content'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(t.Content)

	// Field (4) 'GetCreateAt'
	dst = ssz.MarshalUint64(dst, t.CreateAt)

	// Field (0) 'Type'
	if len(t.Type) > 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, t.Type...)

	// Field (1) 'TaskId'
	if len(t.TaskId) > 128 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, t.TaskId...)

	// Field (2) 'IdentityId'
	if len(t.IdentityId) > 1024 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, t.IdentityId...)

	// Field (3) 'Content'
	if len(t.Content) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, t.Content...)

	return
}

// UnmarshalSSZ ssz unmarshals the TaskEvent object
func (t *TaskEvent) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 24 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2, o3 uint64

	// Offset (0) 'Type'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 24 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'TaskId'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'IdentityId'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Offset (3) 'Content'
	if o3 = ssz.ReadOffset(buf[12:16]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'GetCreateAt'
	t.CreateAt = ssz.UnmarshallUint64(buf[16:24])

	// Field (0) 'Type'
	{
		buf = tail[o0:o1]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(t.Type) == 0 {
			t.Type = make([]byte, 0, len(buf))
		}
		t.Type = append(t.Type, buf...)
	}

	// Field (1) 'TaskId'
	{
		buf = tail[o1:o2]
		if len(buf) > 128 {
			return ssz.ErrBytesLength
		}
		if cap(t.TaskId) == 0 {
			t.TaskId = make([]byte, 0, len(buf))
		}
		t.TaskId = append(t.TaskId, buf...)
	}

	// Field (2) 'IdentityId'
	{
		buf = tail[o2:o3]
		if len(buf) > 1024 {
			return ssz.ErrBytesLength
		}
		if cap(t.IdentityId) == 0 {
			t.IdentityId = make([]byte, 0, len(buf))
		}
		t.IdentityId = append(t.IdentityId, buf...)
	}

	// Field (3) 'Content'
	{
		buf = tail[o3:]
		if len(buf) > 2048 {
			return ssz.ErrBytesLength
		}
		if cap(t.Content) == 0 {
			t.Content = make([]byte, 0, len(buf))
		}
		t.Content = append(t.Content, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the TaskEvent object
func (t *TaskEvent) SizeSSZ() (size int) {
	size = 24

	// Field (0) 'Type'
	size += len(t.Type)

	// Field (1) 'TaskId'
	size += len(t.TaskId)

	// Field (2) 'IdentityId'
	size += len(t.IdentityId)

	// Field (3) 'Content'
	size += len(t.Content)

	return
}

// HashTreeRoot ssz hashes the TaskEvent object
func (t *TaskEvent) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(t)
}

// HashTreeRootWith ssz hashes the TaskEvent object with a hasher
func (t *TaskEvent) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	if len(t.Type) > 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(t.Type)

	// Field (1) 'TaskId'
	if len(t.TaskId) > 128 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(t.TaskId)

	// Field (2) 'IdentityId'
	if len(t.IdentityId) > 1024 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(t.IdentityId)

	// Field (3) 'Content'
	if len(t.Content) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(t.Content)

	// Field (4) 'GetCreateAt'
	hh.PutUint64(t.CreateAt)

	hh.Merkleize(indx)
	return
}
