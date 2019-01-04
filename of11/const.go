/*
 * Copyright (c) 2008 The Board of Trustees of The Leland Stanford Junior University
 * Copyright (c) 2011, 2012 Open Networking Foundation
 * Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler.
 * Copyright 2018, Red Hat, Inc.
 */
// Automatically generated by LOXI from template const.go
// Do not modify

package of11

import (
	"fmt"
	"strings"
)

const (
	// Identifiers from group macro_definitions
	MaxTableNameLen     = 32         // OFP_MAX_TABLE_NAME_LEN
	MaxPortNameLen      = 16         // OFP_MAX_PORT_NAME_LEN
	TCPPort             = 6653       // OFP_TCP_PORT
	SSLPort             = 6653       // OFP_SSL_PORT
	EthAlen             = 6          // OFP_ETH_ALEN
	DefaultMissSendLen  = 128        // OFP_DEFAULT_MISS_SEND_LEN
	OFPFWICMPType       = 64         // OFPFW_ICMP_TYPE
	OFPFWICMPCode       = 128        // OFPFW_ICMP_CODE
	DlTypeEth2Cutoff    = 1536       // OFP_DL_TYPE_ETH2_CUTOFF
	DlTypeNotEthType    = 1535       // OFP_DL_TYPE_NOT_ETH_TYPE
	VLANNone            = 0          // OFP_VLAN_NONE
	OFPMTStandardLength = 88         // OFPMT_STANDARD_LENGTH
	FlowPermanent       = 0          // OFP_FLOW_PERMANENT
	DefaultPriority     = 32768      // OFP_DEFAULT_PRIORITY
	DescStrLen          = 256        // DESC_STR_LEN
	SerialNumLen        = 32         // SERIAL_NUM_LEN
	OFPQAll             = 4294967295 // OFPQ_ALL
	OFPQMinRateUncfg    = 65535      // OFPQ_MIN_RATE_UNCFG
)

const (
	// Identifiers from group nx_action_controller2_prop_type
	Nxac2PtMaxLen       = 0 // NXAC2PT_MAX_LEN
	Nxac2PtControllerID = 1 // NXAC2PT_CONTROLLER_ID
	Nxac2PtReason       = 2 // NXAC2PT_REASON
	Nxac2PtUserdata     = 3 // NXAC2PT_USERDATA
	Nxac2PtPause        = 4 // NXAC2PT_PAUSE
	Nxac2PtMeterID      = 5 // NXAC2PT_METER_ID
)

type NxActionController2PropType uint16

func (self NxActionController2PropType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group nx_bd_algorithms
	NxBdAlgActiveBackup = 0 // NX_BD_ALG_ACTIVE_BACKUP
	NxBdAlgHrw          = 1 // NX_BD_ALG_HRW
)

type NxBdAlgorithms uint16

func (self NxBdAlgorithms) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self NxBdAlgorithms) String() string {
	switch self {
	case NxBdAlgActiveBackup:
		return "\"active_backup\""
	case NxBdAlgHrw:
		return "\"hrw\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for NxBdAlgorithms\"", self)
	}
}

const (
	// Identifiers from group nx_conntrack_flags
	NxCtFCommit = 1 // NX_CT_F_COMMIT
	NxCtFForce  = 2 // NX_CT_F_FORCE
)

type NxConntrackFlags uint16

func (self NxConntrackFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&NxCtFCommit == NxCtFCommit {
		flags = append(flags, "\"Commit\": true")
	}
	if self&NxCtFForce == NxCtFForce {
		flags = append(flags, "\"Force\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group nx_flow_monitor_flags
	NxfmfInitial = 1  // NXFMF_INITIAL
	NxfmfAdd     = 2  // NXFMF_ADD
	NxfmfDelete  = 4  // NXFMF_DELETE
	NxfmfModify  = 8  // NXFMF_MODIFY
	NxfmfActions = 16 // NXFMF_ACTIONS
	NxfmfOwn     = 32 // NXFMF_OWN
)

type NxFlowMonitorFlags uint16

func (self NxFlowMonitorFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&NxfmfInitial == NxfmfInitial {
		flags = append(flags, "\"Initial\": true")
	}
	if self&NxfmfAdd == NxfmfAdd {
		flags = append(flags, "\"Add\": true")
	}
	if self&NxfmfDelete == NxfmfDelete {
		flags = append(flags, "\"Delete\": true")
	}
	if self&NxfmfModify == NxfmfModify {
		flags = append(flags, "\"Modify\": true")
	}
	if self&NxfmfActions == NxfmfActions {
		flags = append(flags, "\"Actions\": true")
	}
	if self&NxfmfOwn == NxfmfOwn {
		flags = append(flags, "\"Own\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group nx_hash_fields
	NxHashFieldsEthSrc           = 0 // NX_HASH_FIELDS_ETH_SRC
	NxHashFieldsSymmetricL4      = 1 // NX_HASH_FIELDS_SYMMETRIC_L4
	NxHashFieldsSymmetricL3L4    = 2 // NX_HASH_FIELDS_SYMMETRIC_L3L4
	NxHashFieldsSymmetricL3L4Udp = 3 // NX_HASH_FIELDS_SYMMETRIC_L3L4_UDP
	NxHashFieldsNwSrc            = 4 // NX_HASH_FIELDS_NW_SRC
	NxHashFieldsNwDst            = 5 // NX_HASH_FIELDS_NW_DST
	NxHashFieldsSymmetricL3      = 6 // NX_HASH_FIELDS_SYMMETRIC_L3
)

type NxHashFields uint16

func (self NxHashFields) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self NxHashFields) String() string {
	switch self {
	case NxHashFieldsEthSrc:
		return "\"eth_src\""
	case NxHashFieldsSymmetricL4:
		return "\"symmetric_l4\""
	case NxHashFieldsSymmetricL3L4:
		return "\"symmetric_l3l4\""
	case NxHashFieldsSymmetricL3L4Udp:
		return "\"symmetric_l3l4_udp\""
	case NxHashFieldsNwSrc:
		return "\"nw_src\""
	case NxHashFieldsNwDst:
		return "\"nw_dst\""
	case NxHashFieldsSymmetricL3:
		return "\"symmetric_l3\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for NxHashFields\"", self)
	}
}

const (
	// Identifiers from group nx_mp_algorithm
	NxMpAlgModuloN       = 0 // NX_MP_ALG_MODULO_N
	NxMpAlgHashThreshold = 1 // NX_MP_ALG_HASH_THRESHOLD
	NxMpAlgHrw           = 2 // NX_MP_ALG_HRW
	NxMpAlgIterHash      = 3 // NX_MP_ALG_ITER_HASH
)

type NxMpAlgorithm uint16

func (self NxMpAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self NxMpAlgorithm) String() string {
	switch self {
	case NxMpAlgModuloN:
		return "\"modulo_n\""
	case NxMpAlgHashThreshold:
		return "\"hash_threshold\""
	case NxMpAlgHrw:
		return "\"hrw\""
	case NxMpAlgIterHash:
		return "\"iter_hash\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for NxMpAlgorithm\"", self)
	}
}

const (
	// Identifiers from group of_action_nx_bundle_slave_type
	NxmOfInPort                  = 2     // NXM_OF_IN_PORT
	NxmOfJustThereToDefinePrefix = 10000 // NXM_OF_JUST_THERE_TO_DEFINE_PREFIX
)

type ActionNxBundleSlaveType uint32

func (self ActionNxBundleSlaveType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self ActionNxBundleSlaveType) String() string {
	switch self {
	case NxmOfInPort:
		return "\"in_port\""
	case NxmOfJustThereToDefinePrefix:
		return "\"just_there_to_define_prefix\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for ActionNxBundleSlaveType\"", self)
	}
}

const (
	// Identifiers from group of_bsn_pdu_slot_num
	BsnPduSlotNumAny = 255 // BSN_PDU_SLOT_NUM_ANY
)

type BsnPduSlotNum uint8

func (self BsnPduSlotNum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group of_nx_nat_range
	NxNatRangeIpv4Min  = 1  // NX_NAT_RANGE_IPV4_MIN
	NxNatRangeIpv4Max  = 2  // NX_NAT_RANGE_IPV4_MAX
	NxNatRangeIpv6Min  = 4  // NX_NAT_RANGE_IPV6_MIN
	NxNatRangeIpv6Max  = 8  // NX_NAT_RANGE_IPV6_MAX
	NxNatRangeProtoMin = 16 // NX_NAT_RANGE_PROTO_MIN
	NxNatRangeProtoMax = 32 // NX_NAT_RANGE_PROTO_MAX
)

type NxNatRange uint16

func (self NxNatRange) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&NxNatRangeIpv4Min == NxNatRangeIpv4Min {
		flags = append(flags, "\"Ipv4Min\": true")
	}
	if self&NxNatRangeIpv4Max == NxNatRangeIpv4Max {
		flags = append(flags, "\"Ipv4Max\": true")
	}
	if self&NxNatRangeIpv6Min == NxNatRangeIpv6Min {
		flags = append(flags, "\"Ipv6Min\": true")
	}
	if self&NxNatRangeIpv6Max == NxNatRangeIpv6Max {
		flags = append(flags, "\"Ipv6Max\": true")
	}
	if self&NxNatRangeProtoMin == NxNatRangeProtoMin {
		flags = append(flags, "\"ProtoMin\": true")
	}
	if self&NxNatRangeProtoMax == NxNatRangeProtoMax {
		flags = append(flags, "\"ProtoMax\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_action_type
	OFPATOutput       = 0     // OFPAT_OUTPUT
	OFPATSetVLANVid   = 1     // OFPAT_SET_VLAN_VID
	OFPATSetVLANPCP   = 2     // OFPAT_SET_VLAN_PCP
	OFPATSetDlSrc     = 3     // OFPAT_SET_DL_SRC
	OFPATSetDlDst     = 4     // OFPAT_SET_DL_DST
	OFPATSetNwSrc     = 5     // OFPAT_SET_NW_SRC
	OFPATSetNwDst     = 6     // OFPAT_SET_NW_DST
	OFPATSetNwTos     = 7     // OFPAT_SET_NW_TOS
	OFPATSetNwEcn     = 8     // OFPAT_SET_NW_ECN
	OFPATSetTpSrc     = 9     // OFPAT_SET_TP_SRC
	OFPATSetTpDst     = 10    // OFPAT_SET_TP_DST
	OFPATCopyTtlOut   = 11    // OFPAT_COPY_TTL_OUT
	OFPATCopyTtlIn    = 12    // OFPAT_COPY_TTL_IN
	OFPATSetMplsLabel = 13    // OFPAT_SET_MPLS_LABEL
	OFPATSetMplsTc    = 14    // OFPAT_SET_MPLS_TC
	OFPATSetMplsTtl   = 15    // OFPAT_SET_MPLS_TTL
	OFPATDecMplsTtl   = 16    // OFPAT_DEC_MPLS_TTL
	OFPATPushVLAN     = 17    // OFPAT_PUSH_VLAN
	OFPATPopVLAN      = 18    // OFPAT_POP_VLAN
	OFPATPushMpls     = 19    // OFPAT_PUSH_MPLS
	OFPATPopMpls      = 20    // OFPAT_POP_MPLS
	OFPATSetQueue     = 21    // OFPAT_SET_QUEUE
	OFPATGroup        = 22    // OFPAT_GROUP
	OFPATSetNwTtl     = 23    // OFPAT_SET_NW_TTL
	OFPATDecNwTtl     = 24    // OFPAT_DEC_NW_TTL
	OFPATExperimenter = 65535 // OFPAT_EXPERIMENTER
)

type ActionType uint16

func (self ActionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bad_action_code
	OFPBACBadType             = 0  // OFPBAC_BAD_TYPE
	OFPBACBadLen              = 1  // OFPBAC_BAD_LEN
	OFPBACBadExperimenter     = 2  // OFPBAC_BAD_EXPERIMENTER
	OFPBACBadExperimenterType = 3  // OFPBAC_BAD_EXPERIMENTER_TYPE
	OFPBACBadOutPort          = 4  // OFPBAC_BAD_OUT_PORT
	OFPBACBadArgument         = 5  // OFPBAC_BAD_ARGUMENT
	OFPBACEperm               = 6  // OFPBAC_EPERM
	OFPBACTooMany             = 7  // OFPBAC_TOO_MANY
	OFPBACBadQueue            = 8  // OFPBAC_BAD_QUEUE
	OFPBACBadOutGroup         = 9  // OFPBAC_BAD_OUT_GROUP
	OFPBACMatchInconsistent   = 10 // OFPBAC_MATCH_INCONSISTENT
	OFPBACUnsupportedOrder    = 11 // OFPBAC_UNSUPPORTED_ORDER
	OFPBACBadTag              = 12 // OFPBAC_BAD_TAG
)

type BadActionCode uint16

func (self BadActionCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bad_instruction_code
	OFPBICUnknownInst       = 0 // OFPBIC_UNKNOWN_INST
	OFPBICUnsupInst         = 1 // OFPBIC_UNSUP_INST
	OFPBICBadTableID        = 2 // OFPBIC_BAD_TABLE_ID
	OFPBICUnsupMetadata     = 3 // OFPBIC_UNSUP_METADATA
	OFPBICUnsupMetadataMask = 4 // OFPBIC_UNSUP_METADATA_MASK
	OFPBICUnsupExpInst      = 5 // OFPBIC_UNSUP_EXP_INST
)

type BadInstructionCode uint16

func (self BadInstructionCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bad_match_code
	OFPBMCBadType       = 0 // OFPBMC_BAD_TYPE
	OFPBMCBadLen        = 1 // OFPBMC_BAD_LEN
	OFPBMCBadTag        = 2 // OFPBMC_BAD_TAG
	OFPBMCBadDlAddrMask = 3 // OFPBMC_BAD_DL_ADDR_MASK
	OFPBMCBadNwAddrMask = 4 // OFPBMC_BAD_NW_ADDR_MASK
	OFPBMCBadWildcards  = 5 // OFPBMC_BAD_WILDCARDS
	OFPBMCBadField      = 6 // OFPBMC_BAD_FIELD
	OFPBMCBadValue      = 7 // OFPBMC_BAD_VALUE
)

type BadMatchCode uint16

func (self BadMatchCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bad_request_code
	OFPBRCBadVersion      = 0 // OFPBRC_BAD_VERSION
	OFPBRCBadType         = 1 // OFPBRC_BAD_TYPE
	OFPBRCBadStat         = 2 // OFPBRC_BAD_STAT
	OFPBRCBadExperimenter = 3 // OFPBRC_BAD_EXPERIMENTER
	OFPBRCBadSubtype      = 4 // OFPBRC_BAD_SUBTYPE
	OFPBRCEperm           = 5 // OFPBRC_EPERM
	OFPBRCBadLen          = 6 // OFPBRC_BAD_LEN
	OFPBRCBufferEmpty     = 7 // OFPBRC_BUFFER_EMPTY
	OFPBRCBufferUnknown   = 8 // OFPBRC_BUFFER_UNKNOWN
	OFPBRCBadTableID      = 9 // OFPBRC_BAD_TABLE_ID
)

type BadRequestCode uint16

func (self BadRequestCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bsn_vport_l2gre_flags
	OFBSNVportL2GreLocalMACIsValid  = 1  // OF_BSN_VPORT_L2GRE_LOCAL_MAC_IS_VALID
	OFBSNVportL2GreDSCPAssign       = 2  // OF_BSN_VPORT_L2GRE_DSCP_ASSIGN
	OFBSNVportL2GreDSCPCopy         = 4  // OF_BSN_VPORT_L2GRE_DSCP_COPY
	OFBSNVportL2GreLoopbackIsValid  = 8  // OF_BSN_VPORT_L2GRE_LOOPBACK_IS_VALID
	OFBSNVportL2GreRateLimitIsValid = 16 // OF_BSN_VPORT_L2GRE_RATE_LIMIT_IS_VALID
)

type BsnVportL2GreFlags uint32

func (self BsnVportL2GreFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFBSNVportL2GreLocalMACIsValid == OFBSNVportL2GreLocalMACIsValid {
		flags = append(flags, "\"LocalMACIsValid\": true")
	}
	if self&OFBSNVportL2GreDSCPAssign == OFBSNVportL2GreDSCPAssign {
		flags = append(flags, "\"DscpAssign\": true")
	}
	if self&OFBSNVportL2GreDSCPCopy == OFBSNVportL2GreDSCPCopy {
		flags = append(flags, "\"DscpCopy\": true")
	}
	if self&OFBSNVportL2GreLoopbackIsValid == OFBSNVportL2GreLoopbackIsValid {
		flags = append(flags, "\"LoopbackIsValid\": true")
	}
	if self&OFBSNVportL2GreRateLimitIsValid == OFBSNVportL2GreRateLimitIsValid {
		flags = append(flags, "\"RateLimitIsValid\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_bsn_vport_q_in_q_untagged
	OFBSNVportQInQUntagged = 65535 // OF_BSN_VPORT_Q_IN_Q_UNTAGGED
)

type BsnVportQInQUntagged uint16

func (self BsnVportQInQUntagged) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_bsn_vport_status
	OFBSNVportStatusOk     = 0 // OF_BSN_VPORT_STATUS_OK
	OFBSNVportStatusFailed = 1 // OF_BSN_VPORT_STATUS_FAILED
)

const (
	// Identifiers from group ofp_capabilities
	OFPCFlowStats  = 1   // OFPC_FLOW_STATS
	OFPCTableStats = 2   // OFPC_TABLE_STATS
	OFPCPortStats  = 4   // OFPC_PORT_STATS
	OFPCGroupStats = 8   // OFPC_GROUP_STATS
	OFPCIpReasm    = 32  // OFPC_IP_REASM
	OFPCQueueStats = 64  // OFPC_QUEUE_STATS
	OFPCARPMatchIp = 128 // OFPC_ARP_MATCH_IP
)

type Capabilities uint32

func (self Capabilities) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPCFlowStats == OFPCFlowStats {
		flags = append(flags, "\"FlowStats\": true")
	}
	if self&OFPCTableStats == OFPCTableStats {
		flags = append(flags, "\"TableStats\": true")
	}
	if self&OFPCPortStats == OFPCPortStats {
		flags = append(flags, "\"PortStats\": true")
	}
	if self&OFPCGroupStats == OFPCGroupStats {
		flags = append(flags, "\"GroupStats\": true")
	}
	if self&OFPCIpReasm == OFPCIpReasm {
		flags = append(flags, "\"IpReasm\": true")
	}
	if self&OFPCQueueStats == OFPCQueueStats {
		flags = append(flags, "\"QueueStats\": true")
	}
	if self&OFPCARPMatchIp == OFPCARPMatchIp {
		flags = append(flags, "\"ArpMatchIp\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_config_flags
	OFPCFragNormal             = 0 // OFPC_FRAG_NORMAL
	OFPCFragDrop               = 1 // OFPC_FRAG_DROP
	OFPCFragReasm              = 2 // OFPC_FRAG_REASM
	OFPCFragMask               = 3 // OFPC_FRAG_MASK
	OFPCInvalidTtlToController = 4 // OFPC_INVALID_TTL_TO_CONTROLLER
)

type ConfigFlags uint16

func (self ConfigFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPCFragNormal == OFPCFragNormal {
		flags = append(flags, "\"FragNormal\": true")
	}
	if self&OFPCFragDrop == OFPCFragDrop {
		flags = append(flags, "\"FragDrop\": true")
	}
	if self&OFPCFragReasm == OFPCFragReasm {
		flags = append(flags, "\"FragReasm\": true")
	}
	if self&OFPCFragMask == OFPCFragMask {
		flags = append(flags, "\"FragMask\": true")
	}
	if self&OFPCInvalidTtlToController == OFPCInvalidTtlToController {
		flags = append(flags, "\"InvalidTtlToController\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_cs_states
	CsNew         = 0  // OFP_CS_NEW
	CsEstablished = 1  // OFP_CS_ESTABLISHED
	CsRelated     = 2  // OFP_CS_RELATED
	CsReplyDir    = 4  // OFP_CS_REPLY_DIR
	CsInvalid     = 8  // OFP_CS_INVALID
	CsTracked     = 16 // OFP_CS_TRACKED
	CsSrcNat      = 32 // OFP_CS_SRC_NAT
	CsDstNat      = 64 // OFP_CS_DST_NAT
)

type CsStates uint32

func (self CsStates) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&CsNew == CsNew {
		flags = append(flags, "\"New\": true")
	}
	if self&CsEstablished == CsEstablished {
		flags = append(flags, "\"Established\": true")
	}
	if self&CsRelated == CsRelated {
		flags = append(flags, "\"Related\": true")
	}
	if self&CsReplyDir == CsReplyDir {
		flags = append(flags, "\"ReplyDir\": true")
	}
	if self&CsInvalid == CsInvalid {
		flags = append(flags, "\"Invalid\": true")
	}
	if self&CsTracked == CsTracked {
		flags = append(flags, "\"Tracked\": true")
	}
	if self&CsSrcNat == CsSrcNat {
		flags = append(flags, "\"SrcNat\": true")
	}
	if self&CsDstNat == CsDstNat {
		flags = append(flags, "\"DstNat\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_ed_nsh_prop_type
	OFPPPTPropNshNone   = 0 // OFPPPT_PROP_NSH_NONE
	OFPPPTPropNshMdtype = 1 // OFPPPT_PROP_NSH_MDTYPE
	OFPPPTPropNshTlv    = 2 // OFPPPT_PROP_NSH_TLV
)

type EdNshPropType uint8

func (self EdNshPropType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_ed_prop_class
	OFPPPCBasic        = 0     // OFPPPC_BASIC
	OFPPPCMpls         = 1     // OFPPPC_MPLS
	OFPPPCGRE          = 2     // OFPPPC_GRE
	OFPPPCGtp          = 3     // OFPPPC_GTP
	OFPPPCNsh          = 4     // OFPPPC_NSH
	OFPPPCExperimenter = 65535 // OFPPPC_EXPERIMENTER
)

type EdPropClass uint16

func (self EdPropClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_error_type
	OFPETHelloFailed        = 0  // OFPET_HELLO_FAILED
	OFPETBadRequest         = 1  // OFPET_BAD_REQUEST
	OFPETBadAction          = 2  // OFPET_BAD_ACTION
	OFPETBadInstruction     = 3  // OFPET_BAD_INSTRUCTION
	OFPETBadMatch           = 4  // OFPET_BAD_MATCH
	OFPETFlowModFailed      = 5  // OFPET_FLOW_MOD_FAILED
	OFPETGroupModFailed     = 6  // OFPET_GROUP_MOD_FAILED
	OFPETPortModFailed      = 7  // OFPET_PORT_MOD_FAILED
	OFPETTableModFailed     = 8  // OFPET_TABLE_MOD_FAILED
	OFPETQueueOpFailed      = 9  // OFPET_QUEUE_OP_FAILED
	OFPETSwitchConfigFailed = 10 // OFPET_SWITCH_CONFIG_FAILED
)

type ErrorType uint16

func (self ErrorType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_flow_mod_command
	OFPFCAdd          = 0 // OFPFC_ADD
	OFPFCModify       = 1 // OFPFC_MODIFY
	OFPFCModifyStrict = 2 // OFPFC_MODIFY_STRICT
	OFPFCDelete       = 3 // OFPFC_DELETE
	OFPFCDeleteStrict = 4 // OFPFC_DELETE_STRICT
)

type FlowModCommand uint8

func (self FlowModCommand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_flow_mod_failed_code
	OFPFMFCUnknown    = 0 // OFPFMFC_UNKNOWN
	OFPFMFCTableFull  = 1 // OFPFMFC_TABLE_FULL
	OFPFMFCBadTableID = 2 // OFPFMFC_BAD_TABLE_ID
	OFPFMFCOverlap    = 3 // OFPFMFC_OVERLAP
	OFPFMFCEperm      = 4 // OFPFMFC_EPERM
	OFPFMFCBadTimeout = 5 // OFPFMFC_BAD_TIMEOUT
	OFPFMFCBadCommand = 6 // OFPFMFC_BAD_COMMAND
)

type FlowModFailedCode uint16

func (self FlowModFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_flow_mod_flags
	OFPFFSendFlowRem  = 1 // OFPFF_SEND_FLOW_REM
	OFPFFCheckOverlap = 2 // OFPFF_CHECK_OVERLAP
)

type FlowModFlags uint16

func (self FlowModFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPFFSendFlowRem == OFPFFSendFlowRem {
		flags = append(flags, "\"SendFlowRem\": true")
	}
	if self&OFPFFCheckOverlap == OFPFFCheckOverlap {
		flags = append(flags, "\"CheckOverlap\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_flow_removed_reason
	OFPRRIdleTimeout = 0 // OFPRR_IDLE_TIMEOUT
	OFPRRHardTimeout = 1 // OFPRR_HARD_TIMEOUT
	OFPRRDelete      = 2 // OFPRR_DELETE
	OFPRRGroupDelete = 3 // OFPRR_GROUP_DELETE
)

type FlowRemovedReason uint8

func (self FlowRemovedReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_flow_wildcards
	OFPFWInPort    = 1    // OFPFW_IN_PORT
	OFPFWDlVLAN    = 2    // OFPFW_DL_VLAN
	OFPFWDlVLANPCP = 4    // OFPFW_DL_VLAN_PCP
	OFPFWDlType    = 8    // OFPFW_DL_TYPE
	OFPFWNwTos     = 16   // OFPFW_NW_TOS
	OFPFWNwProto   = 32   // OFPFW_NW_PROTO
	OFPFWTpSrc     = 64   // OFPFW_TP_SRC
	OFPFWTpDst     = 128  // OFPFW_TP_DST
	OFPFWMplsLabel = 256  // OFPFW_MPLS_LABEL
	OFPFWMplsTc    = 512  // OFPFW_MPLS_TC
	OFPFWAll       = 1023 // OFPFW_ALL
)

type FlowWildcards uint32

func (self FlowWildcards) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPFWInPort == OFPFWInPort {
		flags = append(flags, "\"InPort\": true")
	}
	if self&OFPFWDlVLAN == OFPFWDlVLAN {
		flags = append(flags, "\"DlVLAN\": true")
	}
	if self&OFPFWDlVLANPCP == OFPFWDlVLANPCP {
		flags = append(flags, "\"DlVLANPCP\": true")
	}
	if self&OFPFWDlType == OFPFWDlType {
		flags = append(flags, "\"DlType\": true")
	}
	if self&OFPFWNwTos == OFPFWNwTos {
		flags = append(flags, "\"NwTos\": true")
	}
	if self&OFPFWNwProto == OFPFWNwProto {
		flags = append(flags, "\"NwProto\": true")
	}
	if self&OFPFWTpSrc == OFPFWTpSrc {
		flags = append(flags, "\"TpSrc\": true")
	}
	if self&OFPFWTpDst == OFPFWTpDst {
		flags = append(flags, "\"TpDst\": true")
	}
	if self&OFPFWMplsLabel == OFPFWMplsLabel {
		flags = append(flags, "\"MplsLabel\": true")
	}
	if self&OFPFWMplsTc == OFPFWMplsTc {
		flags = append(flags, "\"MplsTc\": true")
	}
	if self&OFPFWAll == OFPFWAll {
		flags = append(flags, "\"All\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_group
	OFPGMax = 4294967040 // OFPG_MAX
	OFPGAll = 4294967292 // OFPG_ALL
	OFPGAny = 4294967295 // OFPG_ANY
)

type Group uint32

func (self Group) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_group_mod_command
	OFPGCAdd    = 0 // OFPGC_ADD
	OFPGCModify = 1 // OFPGC_MODIFY
	OFPGCDelete = 2 // OFPGC_DELETE
)

type GroupModCommand uint16

func (self GroupModCommand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_group_mod_failed_code
	OFPGMFCGroupExists         = 0 // OFPGMFC_GROUP_EXISTS
	OFPGMFCInvalidGroup        = 1 // OFPGMFC_INVALID_GROUP
	OFPGMFCWeightUnsupported   = 2 // OFPGMFC_WEIGHT_UNSUPPORTED
	OFPGMFCOutOfGroups         = 3 // OFPGMFC_OUT_OF_GROUPS
	OFPGMFCOutOfBuckets        = 4 // OFPGMFC_OUT_OF_BUCKETS
	OFPGMFCChainingUnsupported = 5 // OFPGMFC_CHAINING_UNSUPPORTED
	OFPGMFCWatchUnsupported    = 6 // OFPGMFC_WATCH_UNSUPPORTED
	OFPGMFCLoop                = 7 // OFPGMFC_LOOP
	OFPGMFCUnknownGroup        = 8 // OFPGMFC_UNKNOWN_GROUP
)

type GroupModFailedCode uint16

func (self GroupModFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_group_type
	OFPGTAll      = 0 // OFPGT_ALL
	OFPGTSelect   = 1 // OFPGT_SELECT
	OFPGTIndirect = 2 // OFPGT_INDIRECT
	OFPGTFf       = 3 // OFPGT_FF
)

type GroupType uint8

func (self GroupType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_header_type_namespaces
	OFPHTNOnf        = 0 // OFPHTN_ONF
	OFPHTNEthertype  = 1 // OFPHTN_ETHERTYPE
	OFPHTNIpProto    = 2 // OFPHTN_IP_PROTO
	OFPHTNUdpTCPPort = 3 // OFPHTN_UDP_TCP_PORT
	OFPHTNIpv4Option = 4 // OFPHTN_IPV4_OPTION
)

type HeaderTypeNamespaces uint16

func (self HeaderTypeNamespaces) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self HeaderTypeNamespaces) String() string {
	switch self {
	case OFPHTNOnf:
		return "\"onf\""
	case OFPHTNEthertype:
		return "\"ethertype\""
	case OFPHTNIpProto:
		return "\"ip_proto\""
	case OFPHTNUdpTCPPort:
		return "\"udp_tcp_port\""
	case OFPHTNIpv4Option:
		return "\"ipv4_option\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for HeaderTypeNamespaces\"", self)
	}
}

const (
	// Identifiers from group ofp_hello_failed_code
	OFPHFCIncompatible = 0 // OFPHFC_INCOMPATIBLE
	OFPHFCEperm        = 1 // OFPHFC_EPERM
)

type HelloFailedCode uint16

func (self HelloFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_instruction_type
	OFPITGotoTable     = 1     // OFPIT_GOTO_TABLE
	OFPITWriteMetadata = 2     // OFPIT_WRITE_METADATA
	OFPITWriteActions  = 3     // OFPIT_WRITE_ACTIONS
	OFPITApplyActions  = 4     // OFPIT_APPLY_ACTIONS
	OFPITClearActions  = 5     // OFPIT_CLEAR_ACTIONS
	OFPITExperimenter  = 65535 // OFPIT_EXPERIMENTER
)

type InstructionType uint16

func (self InstructionType) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPITGotoTable == OFPITGotoTable {
		flags = append(flags, "\"GotoTable\": true")
	}
	if self&OFPITWriteMetadata == OFPITWriteMetadata {
		flags = append(flags, "\"WriteMetadata\": true")
	}
	if self&OFPITWriteActions == OFPITWriteActions {
		flags = append(flags, "\"WriteActions\": true")
	}
	if self&OFPITApplyActions == OFPITApplyActions {
		flags = append(flags, "\"ApplyActions\": true")
	}
	if self&OFPITClearActions == OFPITClearActions {
		flags = append(flags, "\"ClearActions\": true")
	}
	if self&OFPITExperimenter == OFPITExperimenter {
		flags = append(flags, "\"Experimenter\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_match_type
	OFPMTStandard = 0 // OFPMT_STANDARD
)

type MatchType uint16

func (self MatchType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_packet_in_reason
	OFPRNoMatch = 0 // OFPR_NO_MATCH
	OFPRAction  = 1 // OFPR_ACTION
)

type PacketInReason uint8

func (self PacketInReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_packet_type
	PtEthernet     = 0          // OFP_PT_ETHERNET
	PtUseNextProto = 65534      // OFP_PT_USE_NEXT_PROTO
	PtIpv4         = 67584      // OFP_PT_IPV4
	PtMpls         = 100423     // OFP_PT_MPLS
	PtMplsMc       = 100424     // OFP_PT_MPLS_MC
	PtNsh          = 100687     // OFP_PT_NSH
	PtUnknown      = 4294967295 // OFP_PT_UNKNOWN
)

type PacketType uint32

func (self PacketType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self PacketType) String() string {
	switch self {
	case PtEthernet:
		return "\"ethernet\""
	case PtUseNextProto:
		return "\"use_next_proto\""
	case PtIpv4:
		return "\"ipv4\""
	case PtMpls:
		return "\"mpls\""
	case PtMplsMc:
		return "\"mpls_mc\""
	case PtNsh:
		return "\"nsh\""
	case PtUnknown:
		return "\"unknown\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for PacketType\"", self)
	}
}

const (
	// Identifiers from group ofp_port
	OFPPMax        = 4294967040 // OFPP_MAX
	OFPPInPort     = 4294967288 // OFPP_IN_PORT
	OFPPTable      = 4294967289 // OFPP_TABLE
	OFPPNormal     = 4294967290 // OFPP_NORMAL
	OFPPFlood      = 4294967291 // OFPP_FLOOD
	OFPPAll        = 4294967292 // OFPP_ALL
	OFPPController = 4294967293 // OFPP_CONTROLLER
	OFPPLocal      = 4294967294 // OFPP_LOCAL
	OFPPAny        = 4294967295 // OFPP_ANY
)

type Port uint32

func (self Port) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_port_config
	OFPPCPortDown      = 1          // OFPPC_PORT_DOWN
	OFPPCNoRecv        = 4          // OFPPC_NO_RECV
	OFPPCNoFwd         = 32         // OFPPC_NO_FWD
	OFPPCNoPacketIn    = 64         // OFPPC_NO_PACKET_IN
	OFPPCBSNMirrorDest = 2147483648 // OFPPC_BSN_MIRROR_DEST
)

type PortConfig uint32

func (self PortConfig) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPCPortDown == OFPPCPortDown {
		flags = append(flags, "\"PortDown\": true")
	}
	if self&OFPPCNoRecv == OFPPCNoRecv {
		flags = append(flags, "\"NoRecv\": true")
	}
	if self&OFPPCNoFwd == OFPPCNoFwd {
		flags = append(flags, "\"NoFwd\": true")
	}
	if self&OFPPCNoPacketIn == OFPPCNoPacketIn {
		flags = append(flags, "\"NoPacketIn\": true")
	}
	if self&OFPPCBSNMirrorDest == OFPPCBSNMirrorDest {
		flags = append(flags, "\"BsnMirrorDest\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_port_features
	OFPPF10MbHd    = 1     // OFPPF_10MB_HD
	OFPPF10MbFd    = 2     // OFPPF_10MB_FD
	OFPPF100MbHd   = 4     // OFPPF_100MB_HD
	OFPPF100MbFd   = 8     // OFPPF_100MB_FD
	OFPPF1GbHd     = 16    // OFPPF_1GB_HD
	OFPPF1GbFd     = 32    // OFPPF_1GB_FD
	OFPPF10GbFd    = 64    // OFPPF_10GB_FD
	OFPPF40GbFd    = 128   // OFPPF_40GB_FD
	OFPPF100GbFd   = 256   // OFPPF_100GB_FD
	OFPPF1TbFd     = 512   // OFPPF_1TB_FD
	OFPPFOther     = 1024  // OFPPF_OTHER
	OFPPFCopper    = 2048  // OFPPF_COPPER
	OFPPFFiber     = 4096  // OFPPF_FIBER
	OFPPFAutoneg   = 8192  // OFPPF_AUTONEG
	OFPPFPause     = 16384 // OFPPF_PAUSE
	OFPPFPauseAsym = 32768 // OFPPF_PAUSE_ASYM
)

type PortFeatures uint32

func (self PortFeatures) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPF10MbHd == OFPPF10MbHd {
		flags = append(flags, "\"10MbHd\": true")
	}
	if self&OFPPF10MbFd == OFPPF10MbFd {
		flags = append(flags, "\"10MbFd\": true")
	}
	if self&OFPPF100MbHd == OFPPF100MbHd {
		flags = append(flags, "\"100MbHd\": true")
	}
	if self&OFPPF100MbFd == OFPPF100MbFd {
		flags = append(flags, "\"100MbFd\": true")
	}
	if self&OFPPF1GbHd == OFPPF1GbHd {
		flags = append(flags, "\"1GbHd\": true")
	}
	if self&OFPPF1GbFd == OFPPF1GbFd {
		flags = append(flags, "\"1GbFd\": true")
	}
	if self&OFPPF10GbFd == OFPPF10GbFd {
		flags = append(flags, "\"10GbFd\": true")
	}
	if self&OFPPF40GbFd == OFPPF40GbFd {
		flags = append(flags, "\"40GbFd\": true")
	}
	if self&OFPPF100GbFd == OFPPF100GbFd {
		flags = append(flags, "\"100GbFd\": true")
	}
	if self&OFPPF1TbFd == OFPPF1TbFd {
		flags = append(flags, "\"1TbFd\": true")
	}
	if self&OFPPFOther == OFPPFOther {
		flags = append(flags, "\"Other\": true")
	}
	if self&OFPPFCopper == OFPPFCopper {
		flags = append(flags, "\"Copper\": true")
	}
	if self&OFPPFFiber == OFPPFFiber {
		flags = append(flags, "\"Fiber\": true")
	}
	if self&OFPPFAutoneg == OFPPFAutoneg {
		flags = append(flags, "\"Autoneg\": true")
	}
	if self&OFPPFPause == OFPPFPause {
		flags = append(flags, "\"Pause\": true")
	}
	if self&OFPPFPauseAsym == OFPPFPauseAsym {
		flags = append(flags, "\"PauseAsym\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_port_mod_failed_code
	OFPPMFCBadPort      = 0 // OFPPMFC_BAD_PORT
	OFPPMFCBadHwAddr    = 1 // OFPPMFC_BAD_HW_ADDR
	OFPPMFCBadConfig    = 2 // OFPPMFC_BAD_CONFIG
	OFPPMFCBadAdvertise = 3 // OFPPMFC_BAD_ADVERTISE
)

type PortModFailedCode uint16

func (self PortModFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_port_reason
	OFPPRAdd    = 0 // OFPPR_ADD
	OFPPRDelete = 1 // OFPPR_DELETE
	OFPPRModify = 2 // OFPPR_MODIFY
)

type PortReason uint8

func (self PortReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_port_state
	OFPPSLinkDown = 1 // OFPPS_LINK_DOWN
	OFPPSBlocked  = 2 // OFPPS_BLOCKED
	OFPPSLive     = 4 // OFPPS_LIVE
)

type PortState uint32

func (self PortState) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPSLinkDown == OFPPSLinkDown {
		flags = append(flags, "\"LinkDown\": true")
	}
	if self&OFPPSBlocked == OFPPSBlocked {
		flags = append(flags, "\"Blocked\": true")
	}
	if self&OFPPSLive == OFPPSLive {
		flags = append(flags, "\"Live\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_queue_op_failed_code
	OFPQOFCBadPort  = 0 // OFPQOFC_BAD_PORT
	OFPQOFCBadQueue = 1 // OFPQOFC_BAD_QUEUE
	OFPQOFCEperm    = 2 // OFPQOFC_EPERM
)

type QueueOpFailedCode uint16

func (self QueueOpFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_queue_properties
	OFPQTNone    = 0 // OFPQT_NONE
	OFPQTMinRate = 1 // OFPQT_MIN_RATE
)

type QueueProperties uint16

func (self QueueProperties) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_stats_reply_flags
	OFPSFReplyMore = 1 // OFPSF_REPLY_MORE
)

type StatsReplyFlags uint16

func (self StatsReplyFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPSFReplyMore == OFPSFReplyMore {
		flags = append(flags, "\"OFPSFReplyMore\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
// Identifiers from group ofp_stats_request_flags
)

type StatsRequestFlags uint16

func (self StatsRequestFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_stats_type
	OFPSTDesc         = 0     // OFPST_DESC
	OFPSTFlow         = 1     // OFPST_FLOW
	OFPSTAggregate    = 2     // OFPST_AGGREGATE
	OFPSTTable        = 3     // OFPST_TABLE
	OFPSTPort         = 4     // OFPST_PORT
	OFPSTQueue        = 5     // OFPST_QUEUE
	OFPSTGroup        = 6     // OFPST_GROUP
	OFPSTGroupDesc    = 7     // OFPST_GROUP_DESC
	OFPSTExperimenter = 65535 // OFPST_EXPERIMENTER
)

type StatsType uint16

func (self StatsType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_switch_config_failed_code
	OFPSCFCBadFlags = 0 // OFPSCFC_BAD_FLAGS
	OFPSCFCBadLen   = 1 // OFPSCFC_BAD_LEN
)

type SwitchConfigFailedCode uint16

func (self SwitchConfigFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_table_config
	OFPTCTableMissController = 0 // OFPTC_TABLE_MISS_CONTROLLER
	OFPTCTableMissContinue   = 1 // OFPTC_TABLE_MISS_CONTINUE
	OFPTCTableMissDrop       = 2 // OFPTC_TABLE_MISS_DROP
	OFPTCTableMissMask       = 3 // OFPTC_TABLE_MISS_MASK
)

type TableConfig uint32

func (self TableConfig) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPTCTableMissController == OFPTCTableMissController {
		flags = append(flags, "\"Controller\": true")
	}
	if self&OFPTCTableMissContinue == OFPTCTableMissContinue {
		flags = append(flags, "\"Continue\": true")
	}
	if self&OFPTCTableMissDrop == OFPTCTableMissDrop {
		flags = append(flags, "\"Drop\": true")
	}
	if self&OFPTCTableMissMask == OFPTCTableMissMask {
		flags = append(flags, "\"Mask\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_table_mod_failed_code
	OFPTMFCBadTable  = 0 // OFPTMFC_BAD_TABLE
	OFPTMFCBadConfig = 1 // OFPTMFC_BAD_CONFIG
)

type TableModFailedCode uint16

func (self TableModFailedCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}

const (
	// Identifiers from group ofp_type
	OFPTHello                 = 0  // OFPT_HELLO
	OFPTError                 = 1  // OFPT_ERROR
	OFPTEchoRequest           = 2  // OFPT_ECHO_REQUEST
	OFPTEchoReply             = 3  // OFPT_ECHO_REPLY
	OFPTExperimenter          = 4  // OFPT_EXPERIMENTER
	OFPTFeaturesRequest       = 5  // OFPT_FEATURES_REQUEST
	OFPTFeaturesReply         = 6  // OFPT_FEATURES_REPLY
	OFPTGetConfigRequest      = 7  // OFPT_GET_CONFIG_REQUEST
	OFPTGetConfigReply        = 8  // OFPT_GET_CONFIG_REPLY
	OFPTSetConfig             = 9  // OFPT_SET_CONFIG
	OFPTPacketIn              = 10 // OFPT_PACKET_IN
	OFPTFlowRemoved           = 11 // OFPT_FLOW_REMOVED
	OFPTPortStatus            = 12 // OFPT_PORT_STATUS
	OFPTPacketOut             = 13 // OFPT_PACKET_OUT
	OFPTFlowMod               = 14 // OFPT_FLOW_MOD
	OFPTGroupMod              = 15 // OFPT_GROUP_MOD
	OFPTPortMod               = 16 // OFPT_PORT_MOD
	OFPTTableMod              = 17 // OFPT_TABLE_MOD
	OFPTStatsRequest          = 18 // OFPT_STATS_REQUEST
	OFPTStatsReply            = 19 // OFPT_STATS_REPLY
	OFPTBarrierRequest        = 20 // OFPT_BARRIER_REQUEST
	OFPTBarrierReply          = 21 // OFPT_BARRIER_REPLY
	OFPTQueueGetConfigRequest = 22 // OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPTQueueGetConfigReply   = 23 // OFPT_QUEUE_GET_CONFIG_REPLY
)

type Type uint8

func (self Type) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", self)), nil
}

func (self Type) String() string {
	switch self {
	case OFPTHello:
		return "\"hello\""
	case OFPTError:
		return "\"error\""
	case OFPTEchoRequest:
		return "\"echo_request\""
	case OFPTEchoReply:
		return "\"echo_reply\""
	case OFPTExperimenter:
		return "\"experimenter\""
	case OFPTFeaturesRequest:
		return "\"features_request\""
	case OFPTFeaturesReply:
		return "\"features_reply\""
	case OFPTGetConfigRequest:
		return "\"get_config_request\""
	case OFPTGetConfigReply:
		return "\"get_config_reply\""
	case OFPTSetConfig:
		return "\"set_config\""
	case OFPTPacketIn:
		return "\"packet_in\""
	case OFPTFlowRemoved:
		return "\"flow_removed\""
	case OFPTPortStatus:
		return "\"port_status\""
	case OFPTPacketOut:
		return "\"packet_out\""
	case OFPTFlowMod:
		return "\"flow_mod\""
	case OFPTGroupMod:
		return "\"group_mod\""
	case OFPTPortMod:
		return "\"port_mod\""
	case OFPTTableMod:
		return "\"table_mod\""
	case OFPTStatsRequest:
		return "\"stats_request\""
	case OFPTStatsReply:
		return "\"stats_reply\""
	case OFPTBarrierRequest:
		return "\"barrier_request\""
	case OFPTBarrierReply:
		return "\"barrier_reply\""
	case OFPTQueueGetConfigRequest:
		return "\"queue_get_config_request\""
	case OFPTQueueGetConfigReply:
		return "\"queue_get_config_reply\""
	default:
		return fmt.Sprintf("\"Invalid value '%d' for Type\"", self)
	}
}

const (
	// Identifiers from group ofp_vlan_id
	OFPVIDAny  = 65534 // OFPVID_ANY
	OFPVIDNone = 65535 // OFPVID_NONE
)

type VlanId uint16

func (self VlanId) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", self)), nil
}
