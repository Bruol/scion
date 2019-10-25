// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type LinkType uint16

// LinkType_TypeID is the unique identifier for the type LinkType.
const LinkType_TypeID = 0x916c98f48c9bbb64

// Values of LinkType.
const (
	LinkType_unset  LinkType = 0
	LinkType_core   LinkType = 1
	LinkType_parent LinkType = 2
	LinkType_child  LinkType = 3
	LinkType_peer   LinkType = 4
)

// String returns the enum's constant name.
func (c LinkType) String() string {
	switch c {
	case LinkType_unset:
		return "unset"
	case LinkType_core:
		return "core"
	case LinkType_parent:
		return "parent"
	case LinkType_child:
		return "child"
	case LinkType_peer:
		return "peer"

	default:
		return ""
	}
}

// LinkTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func LinkTypeFromString(c string) LinkType {
	switch c {
	case "unset":
		return LinkType_unset
	case "core":
		return LinkType_core
	case "parent":
		return LinkType_parent
	case "child":
		return LinkType_child
	case "peer":
		return LinkType_peer

	default:
		return 0
	}
}

type LinkType_List struct{ capnp.List }

func NewLinkType_List(s *capnp.Segment, sz int32) (LinkType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return LinkType_List{l.List}, err
}

func (l LinkType_List) At(i int) LinkType {
	ul := capnp.UInt16List{List: l.List}
	return LinkType(ul.At(i))
}

func (l LinkType_List) Set(i int, v LinkType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type ServiceType uint16

// ServiceType_TypeID is the unique identifier for the type ServiceType.
const ServiceType_TypeID = 0xe2522d291bd06774

// Values of ServiceType.
const (
	ServiceType_unset ServiceType = 0
	ServiceType_bs    ServiceType = 1
	ServiceType_ps    ServiceType = 2
	ServiceType_cs    ServiceType = 3
	ServiceType_sb    ServiceType = 4
	ServiceType_ds    ServiceType = 5
	ServiceType_br    ServiceType = 6
	ServiceType_sig   ServiceType = 7
	ServiceType_hps   ServiceType = 8
)

// String returns the enum's constant name.
func (c ServiceType) String() string {
	switch c {
	case ServiceType_unset:
		return "unset"
	case ServiceType_bs:
		return "bs"
	case ServiceType_ps:
		return "ps"
	case ServiceType_cs:
		return "cs"
	case ServiceType_sb:
		return "sb"
	case ServiceType_ds:
		return "ds"
	case ServiceType_br:
		return "br"
	case ServiceType_sig:
		return "sig"
	case ServiceType_hps:
		return "hps"

	default:
		return ""
	}
}

// ServiceTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ServiceTypeFromString(c string) ServiceType {
	switch c {
	case "unset":
		return ServiceType_unset
	case "bs":
		return ServiceType_bs
	case "ps":
		return ServiceType_ps
	case "cs":
		return ServiceType_cs
	case "sb":
		return ServiceType_sb
	case "ds":
		return ServiceType_ds
	case "br":
		return ServiceType_br
	case "sig":
		return ServiceType_sig
	case "hps":
		return ServiceType_hps

	default:
		return 0
	}
}

type ServiceType_List struct{ capnp.List }

func NewServiceType_List(s *capnp.Segment, sz int32) (ServiceType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ServiceType_List{l.List}, err
}

func (l ServiceType_List) At(i int) ServiceType {
	ul := capnp.UInt16List{List: l.List}
	return ServiceType(ul.At(i))
}

func (l ServiceType_List) Set(i int, v ServiceType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_fa01960eced2b529 = "x\xda\x12\x88u`\x12d\xdd\xce\xc0\x10\xc8\xc1\xca\xf6" +
	"?e\xf7\xec\x9e/3r&2\x08\xf22\xfd\xd7\xdc" +
	"z\xe9\x1c\xdf4\xc6_\x0c\x0c\x8c\x82\x8e\x9b\x04=\xd9" +
	"\x19\x18\x04]\xeb\x19\x18\xff\x97\xa4_\x90\xd6\xd4\x0dz" +
	"\x84\xa1\xaa\xf2\x94`+HU\xe3}\x06\xc6\xff\xc9\xf9" +
	"\xb9\xb9\xf9yz\xc9\x8c\x89\x05y\x05V>\x99y\xf2" +
	"\xd9!\x95\x05\xa9\x01\x8c\x8c\x81\"\x8cL\x0c\x0c\x82\xa6" +
	"F\x0c\x0c\x8c\x8c\x82\xbaZ\x0c\x0c\x8cL\x82\xaaV\x0c" +
	"\x0c\x8c\xcc\x82\xb2 A\x16AQ-\x06\x06\xf9\xd2\xbc" +
	"\xe2\xd4\x12\xfe\xe4\xfc\xa2T\xfb\x82\xc4\xa2\xd4\xbc\x12\xf9" +
	"\xe4\x8c\xcc\x9c\x14\xfe\x82\xd4\xd4\"\xb8\xf1L`\xe3\x83" +
	"S\x8b\xca2\x93SA\x1600\x80\xacP\x01[\x91" +
	"\x0a\xb1\"V\x0alE\xa8\x14\xd8\x0a_)\xb0\x15\xae" +
	" \x8aU\xd0\x16D\xb1\x09\x9a\x82(vA]%\x06" +
	"\x06F\x0eAU%\x98\xf5\xccI\xc5\xcc\x05\xc5\xcc\xc9" +
	"\xc5\xcc\xc5I\xcc)\xc5\xccIE\xec\xc5\x99\xe9\xec\x19" +
	"\x05\xc5\x80\x00\x00\x00\xff\xffj\xfcLl"

func init() {
	schemas.Register(schema_fa01960eced2b529,
		0x916c98f48c9bbb64,
		0xe2522d291bd06774)
}