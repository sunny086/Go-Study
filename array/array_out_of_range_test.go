package array

import "testing"

func TestArrayOutOfRange(t *testing.T) {
	var ProtocolTuple = []string{"equipment", "modbus", "opcda", "iec104", "dnp3", "fins", "mms", "s7", "profnet_cm", "goose", "sv", "enip", "opcua", "pnrt_dcp", "http", "ftp", "bacnet", "tls", "ssh", "imap", "smb", "smb2", "network", "pn_rt", "baosight", "dns", "netbios", "smtp", "telnet", "pop3", "snmp", "egd", "scnet", "scnet2", "unknown"}
	t.Log(len(ProtocolTuple))
	s := ProtocolTuple[111]
	t.Log(s)
}
