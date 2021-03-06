// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/hashtree/hashtree.proto

/*
	Package hashtree is a generated protocol buffer package.

	It is generated from these files:
		server/pkg/hashtree/hashtree.proto

	It has these top-level messages:
		FileNodeProto
		DirectoryNodeProto
		NodeProto
		HashTreeProto
*/
package hashtree

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// FileNodeProto is a node corresponding to a file (which is also a leaf node).
type FileNodeProto struct {
	// Object references an object in the object store which contains the content
	// of the data.
	Objects []*pfs.Object `protobuf:"bytes,4,rep,name=objects" json:"objects,omitempty"`
}

func (m *FileNodeProto) Reset()                    { *m = FileNodeProto{} }
func (m *FileNodeProto) String() string            { return proto.CompactTextString(m) }
func (*FileNodeProto) ProtoMessage()               {}
func (*FileNodeProto) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{0} }

func (m *FileNodeProto) GetObjects() []*pfs.Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

// DirectoryNodeProto is a node corresponding to a directory.
type DirectoryNodeProto struct {
	// Children of this directory. Note that paths are relative, so if "/foo/bar"
	// has a child "baz", that means that there is a file at "/foo/bar/baz".
	//
	// 'Children' is ordered alphabetically, to quickly check if a new file is
	// overwriting an existing one.
	Children []string `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
}

func (m *DirectoryNodeProto) Reset()                    { *m = DirectoryNodeProto{} }
func (m *DirectoryNodeProto) String() string            { return proto.CompactTextString(m) }
func (*DirectoryNodeProto) ProtoMessage()               {}
func (*DirectoryNodeProto) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{1} }

func (m *DirectoryNodeProto) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

// NodeProto is a node in the file tree (either a file or a directory)
type NodeProto struct {
	// Name is the name (not path) of the file/directory (e.g. /lib).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Hash is a hash of the node's name and contents (which includes the
	// BlockRefs of a file and the Children of a directory). This can be used to
	// detect if the name or contents have changed between versions.
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// subtree_size is the of the subtree under node; i.e. if this is a directory,
	// subtree_size includes all children.
	SubtreeSize int64 `protobuf:"varint,3,opt,name=subtree_size,json=subtreeSize,proto3" json:"subtree_size,omitempty"`
	// Exactly one of the following fields must be set. The type of this node will
	// be determined by which field is set.
	FileNode *FileNodeProto      `protobuf:"bytes,4,opt,name=file_node,json=fileNode" json:"file_node,omitempty"`
	DirNode  *DirectoryNodeProto `protobuf:"bytes,5,opt,name=dir_node,json=dirNode" json:"dir_node,omitempty"`
}

func (m *NodeProto) Reset()                    { *m = NodeProto{} }
func (m *NodeProto) String() string            { return proto.CompactTextString(m) }
func (*NodeProto) ProtoMessage()               {}
func (*NodeProto) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{2} }

func (m *NodeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeProto) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *NodeProto) GetSubtreeSize() int64 {
	if m != nil {
		return m.SubtreeSize
	}
	return 0
}

func (m *NodeProto) GetFileNode() *FileNodeProto {
	if m != nil {
		return m.FileNode
	}
	return nil
}

func (m *NodeProto) GetDirNode() *DirectoryNodeProto {
	if m != nil {
		return m.DirNode
	}
	return nil
}

// HashTreeProto is a tree corresponding to the complete file contents of a
// pachyderm repo at a given commit (based on a Merkle Tree). We store one
// HashTree for every PFS commit.
type HashTreeProto struct {
	// Version is an arbitrary version number, set by the corresponding library
	// in hashtree.go.  This ensures that if the hash function used to create
	// these trees is changed, we won't run into errors when deserializing old
	// trees. The current version is 1.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Fs maps each node's path to the NodeProto with that node's details.
	// See "Potential Optimizations" at the end for a compression scheme that
	// could be useful if this map gets too large.
	//
	// Note that the key must end in "/" if an only if the value has .dir_node set
	// (i.e. iff the path points to a directory).
	Fs map[string]*NodeProto `protobuf:"bytes,2,rep,name=fs" json:"fs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HashTreeProto) Reset()                    { *m = HashTreeProto{} }
func (m *HashTreeProto) String() string            { return proto.CompactTextString(m) }
func (*HashTreeProto) ProtoMessage()               {}
func (*HashTreeProto) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{3} }

func (m *HashTreeProto) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HashTreeProto) GetFs() map[string]*NodeProto {
	if m != nil {
		return m.Fs
	}
	return nil
}

func init() {
	proto.RegisterType((*FileNodeProto)(nil), "FileNodeProto")
	proto.RegisterType((*DirectoryNodeProto)(nil), "DirectoryNodeProto")
	proto.RegisterType((*NodeProto)(nil), "NodeProto")
	proto.RegisterType((*HashTreeProto)(nil), "HashTreeProto")
}
func (m *FileNodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileNodeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Objects) > 0 {
		for _, msg := range m.Objects {
			dAtA[i] = 0x22
			i++
			i = encodeVarintHashtree(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DirectoryNodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectoryNodeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, s := range m.Children {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *NodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.SubtreeSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(m.SubtreeSize))
	}
	if m.FileNode != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(m.FileNode.Size()))
		n1, err := m.FileNode.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DirNode != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(m.DirNode.Size()))
		n2, err := m.DirNode.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *HashTreeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashTreeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHashtree(dAtA, i, uint64(m.Version))
	}
	if len(m.Fs) > 0 {
		for k, _ := range m.Fs {
			dAtA[i] = 0x12
			i++
			v := m.Fs[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovHashtree(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovHashtree(uint64(len(k))) + msgSize
			i = encodeVarintHashtree(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintHashtree(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintHashtree(dAtA, i, uint64(v.Size()))
				n3, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n3
			}
		}
	}
	return i, nil
}

func encodeVarintHashtree(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FileNodeProto) Size() (n int) {
	var l int
	_ = l
	if len(m.Objects) > 0 {
		for _, e := range m.Objects {
			l = e.Size()
			n += 1 + l + sovHashtree(uint64(l))
		}
	}
	return n
}

func (m *DirectoryNodeProto) Size() (n int) {
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, s := range m.Children {
			l = len(s)
			n += 1 + l + sovHashtree(uint64(l))
		}
	}
	return n
}

func (m *NodeProto) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHashtree(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.SubtreeSize != 0 {
		n += 1 + sovHashtree(uint64(m.SubtreeSize))
	}
	if m.FileNode != nil {
		l = m.FileNode.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.DirNode != nil {
		l = m.DirNode.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	return n
}

func (m *HashTreeProto) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovHashtree(uint64(m.Version))
	}
	if len(m.Fs) > 0 {
		for k, v := range m.Fs {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHashtree(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHashtree(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHashtree(uint64(mapEntrySize))
		}
	}
	return n
}

func sovHashtree(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHashtree(x uint64) (n int) {
	return sovHashtree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileNodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileNodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileNodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Objects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Objects = append(m.Objects, &pfs.Object{})
			if err := m.Objects[len(m.Objects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DirectoryNodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DirectoryNodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectoryNodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubtreeSize", wireType)
			}
			m.SubtreeSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubtreeSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileNode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileNode == nil {
				m.FileNode = &FileNodeProto{}
			}
			if err := m.FileNode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DirNode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DirNode == nil {
				m.DirNode = &DirectoryNodeProto{}
			}
			if err := m.DirNode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashTreeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HashTreeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashTreeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fs == nil {
				m.Fs = make(map[string]*NodeProto)
			}
			var mapkey string
			var mapvalue *NodeProto
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHashtree
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHashtree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHashtree
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHashtree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHashtree
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthHashtree
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &NodeProto{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHashtree(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHashtree
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fs[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHashtree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHashtree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHashtree
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHashtree(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHashtree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHashtree   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/pkg/hashtree/hashtree.proto", fileDescriptorHashtree) }

var fileDescriptorHashtree = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x75, 0x92, 0xd6, 0x34, 0x37, 0xad, 0x94, 0x51, 0x64, 0xe8, 0x22, 0xc4, 0x80, 0x12, 0x10,
	0xa6, 0x52, 0x41, 0xc4, 0x9d, 0xa2, 0xc5, 0x95, 0xca, 0xe8, 0xbe, 0xa4, 0xc9, 0x8d, 0x19, 0x1b,
	0x93, 0x32, 0x93, 0x16, 0xda, 0xef, 0x70, 0xe1, 0x7f, 0xf8, 0x13, 0x2e, 0xfd, 0x04, 0xe9, 0xfb,
	0x91, 0xc7, 0x4c, 0xd3, 0x96, 0xf2, 0x16, 0x03, 0xe7, 0x9c, 0x7b, 0x2e, 0x9c, 0x7b, 0x06, 0x62,
	0x8d, 0x6a, 0x8b, 0x6a, 0xba, 0x5e, 0x7d, 0x9f, 0x96, 0xa9, 0x2e, 0x5b, 0x85, 0x78, 0x06, 0x7c,
	0xad, 0x9a, 0xb6, 0x99, 0x3c, 0xca, 0x2a, 0x89, 0x75, 0x3b, 0x5d, 0x17, 0xda, 0xbc, 0xa3, 0x1a,
	0xbf, 0x82, 0xd1, 0x5c, 0x56, 0xf8, 0xa9, 0xc9, 0xf1, 0x8b, 0x11, 0xe8, 0x53, 0xf0, 0x9a, 0xe5,
	0x0f, 0xcc, 0x5a, 0xcd, 0x7a, 0x91, 0x9b, 0x04, 0xb3, 0x80, 0x1b, 0xf7, 0x67, 0xab, 0x89, 0xd3,
	0x2c, 0x7e, 0x01, 0xf4, 0xbd, 0x54, 0x98, 0xb5, 0x8d, 0xda, 0x5d, 0x96, 0x27, 0x30, 0xc8, 0x4a,
	0x59, 0xe5, 0x0a, 0x6b, 0xe6, 0x46, 0x6e, 0xe2, 0x8b, 0x33, 0x8f, 0xff, 0x10, 0xf0, 0x2f, 0x4e,
	0x0a, 0xbd, 0x3a, 0xfd, 0x89, 0x8c, 0x44, 0x24, 0xf1, 0x85, 0xc5, 0x46, 0x33, 0x99, 0x99, 0x13,
	0x91, 0x64, 0x28, 0x2c, 0xa6, 0x4f, 0x60, 0xa8, 0x37, 0x4b, 0x73, 0xc6, 0x42, 0xcb, 0x3d, 0x32,
	0x37, 0x22, 0x89, 0x2b, 0x82, 0x4e, 0xfb, 0x2a, 0xf7, 0x48, 0x9f, 0x83, 0x5f, 0xc8, 0x0a, 0x17,
	0x75, 0x93, 0x23, 0xeb, 0x45, 0x24, 0x09, 0x66, 0x0f, 0xf8, 0xd5, 0x51, 0x62, 0x50, 0x74, 0x94,
	0x72, 0x18, 0xe4, 0x52, 0x1d, 0xbd, 0x7d, 0xeb, 0x7d, 0xc8, 0xef, 0x1e, 0x22, 0xbc, 0x5c, 0x2a,
	0xc3, 0xe2, 0x5f, 0x04, 0x46, 0x1f, 0x53, 0x5d, 0x7e, 0x53, 0xd8, 0x25, 0x67, 0xe0, 0x6d, 0x51,
	0x69, 0xd9, 0xd4, 0x36, 0x7c, 0x5f, 0x9c, 0x28, 0x7d, 0x06, 0x4e, 0xa1, 0x99, 0x63, 0x5b, 0x7b,
	0xcc, 0xaf, 0xb6, 0xf8, 0x5c, 0x7f, 0xa8, 0x5b, 0xb5, 0x13, 0x4e, 0xa1, 0x27, 0x6f, 0xc1, 0xeb,
	0x28, 0x1d, 0x83, 0xbb, 0xc2, 0x5d, 0xd7, 0x82, 0x81, 0x34, 0x82, 0xfe, 0x36, 0xad, 0x36, 0x68,
	0x5b, 0x08, 0x66, 0xc0, 0x2f, 0xa1, 0x8e, 0x83, 0x37, 0xce, 0x6b, 0xf2, 0x6e, 0xfc, 0xf7, 0x10,
	0x92, 0x7f, 0x87, 0x90, 0xfc, 0x3f, 0x84, 0xe4, 0xf7, 0x4d, 0x78, 0x6f, 0x79, 0xdf, 0xfe, 0xe7,
	0xcb, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x83, 0xf2, 0x7a, 0x0b, 0x02, 0x00, 0x00,
}
