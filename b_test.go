package b

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/bitcoinschema/go-bob"
)

func TestBitFsURL(t *testing.T) {
	txid := "6ce94f75b88a6c24815d480437f4f06ae895afdab8039ddec10748660c29f910" // donkey kong gif

	bitURL := BitFsURL(txid, 0, 3)

	if bitURL != "https://x.bitfs.network/6ce94f75b88a6c24815d480437f4f06ae895afdab8039ddec10748660c29f910.out.0.3" {
		t.Fatalf("failed url: %s", bitURL)
	}
}

func TestDataURI(t *testing.T) {

	bobData := bob.Tx{}
	bobString := `{ "_id": "5ecd0ef5953f1516596389c4", "tx": { 	"h": "10afc796d06fec11a4b6077012a1522355c82e5de316f4dd5c42ddccd6d61cdb" }, "in": [ 	{ "i": 0, "tape": [ { "cell": [ { "b": "MEUCIQD1JT/xFjFho5At341RRQhD7681Dv3P8lbpXdg/yGaj9gIgWtnk8iuMgdvSNe19nkRlGVqJAnRmefRI8+LIvqC4jORB", "s": "0E\u0002!\u0000�%?�\u00161a��-ߍQE\bC�5\u000e���V�]�?�f��\u0002 Z���+����5�}�De\u0019Z�\u0002tfy�H��Ⱦ����A", "ii": 0, "i": 0 }, { "b": "A6CwrVQl+z6rW/GGBsmRAHTzxxWJ8Dr6Cg7SG9eW2bJc", "s": "\u0003���T%�>�[�\u0006ɑ\u0000t��\u0015��:�\n\u000e�\u001bזٲ\\", "ii": 1, "i": 1 } ], "i": 0 } ], "e": { "h": "8e444af4f3c5c9ec177aeb31b446daf4155c82f48211e1b005c31877dacd68fd", "i": 4, "a": "1HbTsfUXojpWZZ5oZP4h1Tk5riWgqHYDfj" }, "seq": 4294967295 	} ], "out": [ 	{ "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "ii": 0, "i": 0 }, { "op": 106, "ops": "OP_RETURN", "ii": 1, "i": 1 } ], "i": 0 }, { "cell": [ { "b": "MTlIeGlnVjRReUJ2M3RIcFFWY1VFUXlxMXB6WlZkb0F1dA==", "s": "19HxigV4QyBv3tHpQVcUEQyq1pzZVdoAut", "ii": 2, "i": 0 }, { "ii": 3, "i": 1, "ls": "GIF89aX\u0000\u001f\u0000�\u0000\u0000�����x��q��q��j��q��i��i��b��b��c��Z��Z��a��Z��S��Z��S��Z�L��R�KֽRѺR۸K�DֵKϲJڰCޫ<֮CϫC̰JǮJۣ5Ԧ<ʨCΣ<ťCؙ.��C֜)ĝ;̙3��;Ԑ&˒-Ė4ӏ&Ŕ)��4͌&Ɖ&Ō!ͅ\u001e��,��3��,Ɓ\u001e��%�}\u0017��,��!�\u001e��,�z\u0017�|\u001e�t\u0010�y\u001e�{$�r\u000f�t\u0016�|+�s!�w$�q\u0016�s\u001d�k\b�l\u000f�u$�g\b�n\u0016�i\u000f�r$�l\u0016�n\u001d�e\b�i\u0015�m\u001d�b\u0007�j\u001c�_\u0007�e\u000e�f\u0015�a\u000e�b\u0015�` + "`" + `\u0015�\\\u0007�^\u000e�Y\u0007�[\u000e�]\u0015�X\u000e�R\b�U\r�R\b�R\r�J\b�K\u0006�H\u0005�E\u0005{C\u0005k9\bc9\bf3\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000!�\u000bNETSCAPE2.0\u0003\u0001\u0000\u0000\u0000!�\u0010MM Fireworks 1.0\u0000!�\u0004\u0005P\u0000s\u0000,\u0000\u0000\u0000\u0000X\u0000\u001f\u0000\u0000\u0007��pig�����������������M������������i��������i��--C�6������<\u0000\u0002NVV�C\u0000\u0004����M��\"\u0000+aaYP�'\u00004[[Y��ɨ\u0013\u0000c�[ћ\u000f\u0002agg�ٵۧ\b\bglb\f\u0015�\u0002\u0011����\u0001\u0011i��\u0000\u0000 \u000b4L�V�I3F�>T�0�� �\u0000xXȅ��1\r�\u0000�ca�zMx�z��U\u000b\u00180\u001f\"#�iB�\n4h�@@͊\r\u0000\u001fؤyf\u000e]�\u000e�\u0016���\"\u0003�\u0000\u001ch����\u0001�,��El2A@�-a@�I�#\f\r\u00007�>�\u0002�\f\u0001��mr\u00072��\u0019$A�����*��~�\b{F�+gV\u0000����\u0007w�*\u0010\b�\u0000��3$\u0002\u0000p\u0001\u0005���!�\u0016as&\f�\u0007\u001d\b�̰` + "`" + `��ʧSg�\u0014��s�\rn\u0001�\u0001�h&R\u00000�\u0012\u0001A\u0016\u0017\u0000�$p@��\u0015\u0000\t��sǆMh�B=\b�����#\u0000\u0010x0���\u0000kԶ\u0014�\u0014�� ���\u0018䀽4P&\u0000^����L�\u0004\u0007�A\u0006\u0000$0�\u0010\u0002\u00014��\b\u0010P�\u0000\u0015�A�bc�\u0001�\u0006\u0006�q�\u0000\u001a��\u000fc��I{q\u0019�\t}]���\u0007&\u0000�� �\u0010��\u001a\u0010@` + "`" + `��� ` + "`" + `@\u0011J\u0014��\u000e%` + "`" + `'�p\u001f@�\u0001h\u0000\f` + "`" + `B�\u001bh\u0000\u0000�\tF��%\u00119%�u�5�\u0001\u000fN\u0005t�\u0010@5\u0000@\u001b��\u0004]\b\u0007<Q\u001d��PG�9\u0003)�F\u0011�-d'\u0017\u0004��\u0012\u0000\\Y�9v�\b�%A<��\u0019\u000e� \u0000\u0014�yy�\u000eb���\u00194\u0014��\u0004\u0000�i�z�U\u0007�\u000e\u0000�IgZBi0�Y\u000e��\u0006�9T���\"�7�%�\u0011\u0018�/'\u0010 B\u0016\f\u0000��Y>\\\u0000\u0000\u0012mxA�\u0011��Q@\u0003B�t��^s�hCRP�\n%lq�\f�V{m�n]�U\u00167hF\u0010A�\u0005\u0000\u000e8�1�bO<��\u0000+8\u0001�\u0018mh\u0011\u0000\u0010C���\t��PEC0<\u0000�\u0003J��\u0001\u0000\u0017\\` + "`" + `�\u0014ںH�\u0018\u0018\u0000` + "`" + `��\u0005\u000f[��L�h\u0002�\u0016i\u0000a�\t\u001f���9\u001f��ġ+` + "`" + ` \u0006\u001be��B\u0015gD�A@l�\u0010\u0000\u0019c\u0014�T\u0006,�h\u0006\u0019x@�\t��` + "`" + `\u0002\u000bb���\u0007��\u000b$0���[5a�9k�Y�7�c��c�\u0006�\u0006\t\u0018uI\u0016��\bE\u0018\u00011T�B�YSvC�\u00112H�(q�\twe�V��z��LXb\u001d��\u0006L<�v�x\u0003\u000e(Xb]�\u001b4��v5ֳfbE\u0016�h�AX�b:饛\u000e��M�\u0007z�[S:�V���%�\u001b4��\u001d\u0012�1�]S����O�2\u0015{� ���8\u000f���HO���ک���w�����/�\u001f�q�o���/��엿��������o���\u000f \u0000!�\u0004\u00052\u0000s\u0000,\b\u0000\u0003\u0000J\u0000\u0017\u0000\u0000\u0007��M�����������������������������<\u0019\"0��0\u00190�C--�0<��-\u0001\u0019\u0019\u001c\u000f\u0000\u000f\u001c\u001c\f\u001cYVV��\u0013\u0000\b�\u0004\u0015���-\b[aa;\u0000;cc;\u001ea[��F\u0000\u001e�;\t��6����Va�\u0000Eligc7\f\u0001\u0001:PP�P\u0000\u001fllb^Բ\u0018Y���\u0015\u0017\u0019\u0010d\b\"�A���l̺7Ɋ�04ع\u001bS\"\u0000\u0017.\u0001\u0018lɂoо~/^\u0004\\��\n\u0002\u0004�\\\u0000�����#/�\b\u00110sd�H����xFJ\u0000\f�\u001c\u0004���\u0010�\u0004&L\u0018�\u0016f�E\u0003\tҤ�V\"뻧b\u0012\u0010���\u0012\f\u0000>؜ɸ��\u0019\u0012\u0000��m�\t��\u0012%\u0014>�\u0001ydK\u0014\u0000\u001b�E\u001b���Z\u0000$ �8�X�OIgӞ� M뇸l�hkr�M�\"g> PG!\u0000�~7\u0000\u0018p�&�j\u0010m�\u0000(�&��H��RN���W�&���Be�\u0018&\u0014\f�fsY\u0005�4[\u0000\u0004P�F�F�!�\f@��\u0001/�\u0002\u0017��L\t�\u001a�\u0018�\u0002�G�4���\u001e��n�68�n�d��-\u0007E0��\u0000@X��\u0018Z�&H\u000b�\u0015Q�\r\u0002�\u0010�t@��\u0014\u000ei|0@\u001a�1�\r�m\u0000�\u0005J�1�$\"��\u0001\t&�0\u0006\u0011,.��\u0018$�R�\u0007Q}��\u00127�PF\u001b_Ȑ�;;��N\u0016Ap@\u0005vjD��։�d��Fc���;cdC�\u0016N�E���i��\u0019�T\u0019\u00060R�Uf\u001aYf�\b\u0014Y�1�3�P��fM@�E���9�\u0018t��E\u0016[���3#\u0011*�%V\u0010Z蠍\u000eJ�!�Jj)�Aݣ� �vJ�%�6��ʔj*\"�\u0000\u0000!�\u0004\u0005P\u0000s\u0000,\u0010\u0000\u0007\u0000:\u0000\u0016\u0000\u0000\u0007��M���������C\u000b6�����\u0001\u0000V��6--<�\"�����\u0000\bac[P�\"\b\u0000\u0004Y��\u0000\u0000\u0015+���\u0000\u0018iiaV�\u0019�TaY�'\u0002\u0000QccĈ����\u0000%llc��\u0019\u0018\u0001\u0017��\u000f\u0015\u0000Elg�\u0019\u000b�F\"\u0013\u0002\u001dPP\u000f\u0000��'\u0004.A�'PA�\"�\u0000;lx��V�D\u00005g�\u0018\t�\u0004\u0000\u000e7i�,@` + "`" + `��4!\u0016il�@i��\u0000+�x@@c\u0004�\bY\u0010\u0010 \u0000��\u0011\u0000Tv�*D�\u000b\u0000\u0014�DT�r\u0013�E\u0017[�\u0004pp�\u0003\u0000\u001f/(x\u0000p�E�1\u0014\u0012�\u0019� ��\u0000\\n�+�4\r9C$�X\bP���0\u0006 �9�tə3\t�\u0012�!\u0001�ʁ4Fـ\b�a\u0003�\"��\u001e\u0000���'\u0003v-\u0003K�M\u001cW\u0006��\u001b0\u0001�1�(�\u0000Р\f��m\u001a\u001c0Cm�\u0019/���Q\u0002\u0000��\u0006�z\u0019\n DY�\u0001\u001bT��\u0000�2�4�\u0006|H�ل��rESc��)\u0007\u0000��a�\u0006\u000b���Fm��C�\u0017�F3\u0003�\n�4\u0010Z\u001f\u001f \u0017Dk6b\flْe\u0005�'nXH�K��@A6�\u0000�\u0007Lؖ@\u001bh\u0014�\u0000\u0000$T!\u0003gJ\u0018P\u0002f\t\u001c�F\u001bf\u00000�\t` + "`" + `\u0004P�\t\u0003�0_\u0016N\u0004�C\u001b\u0016�FI\u0011�dׄ\b\u0019p��\fc��C\u0011m|���&��F\u000e&���[M�@�\u0018+xAͭ\u0017A\u0016��\u00074�'N\u0000W���\n�\u0019\u0012�Wa\u0004��)�\u0005�˘ʘ�V�a���\u0018\u0000��EP*P@$��\u0011\u0002EP�laś�m\u0011Ɨ�\u0004:L\u0016Y�1h����LwK\u0018J�\u0000^\u0000�\f3��I�\u0015�d\n\u0005��vJ����éE>Z[���WH��hz\u000b!6\u0000�\u0001\u0004\t\u0010�g/.���!4�P�TeR�박���0�\u0012�즉ꙫ��Zb\u000b\"�\u0000\u0000;\u0000", "lb": "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" }, { "b": "aW1hZ2UvZ2lm", "s": "image/gif", "ii": 4, "i": 2 }, { "b": "YmluYXJ5", "s": "binary", "ii": 5, "i": 3 }, { "b": "dHdldGNoX3R3ZW1iZWQxNTg4NjI4NjMxMzY0LmdpZg==", "s": "twetch_twembed1588628631364.gif", "ii": 6, "i": 4 } ], "i": 1 }, { "cell": [ { "b": "MVB1UWE3SzYyTWlLQ3Rzc1NMS3kxa2g1NldXVTdNdFVSNQ==", "s": "1PuQa7K62MiKCtssSLKy1kh56WWU7MtUR5", "ii": 8, "i": 0 }, { "b": "U0VU", "s": "SET", "ii": 9, "i": 1 }, { "b": "dHdkYXRhX2pzb24=", "s": "twdata_json", "ii": 10, "i": 2 }, { "b": "bnVsbA==", "s": "null", "ii": 11, "i": 3 }, { "b": "dXJs", "s": "url", "ii": 12, "i": 4 }, { "b": "bnVsbA==", "s": "null", "ii": 13, "i": 5 }, { "b": "Y29tbWVudA==", "s": "comment", "ii": 14, "i": 6 }, { "b": "V2hhdCdzIHRoZSBmaXJzdCBzdGVwIHRvd2FyZHMgdW5pdmVyc2FsIGJhc2ljIGluY29tZT8KCiQxMjAwISA=", "s": "What's the first step towards universal basic income?\n\n$1200! ", "ii": 15, "i": 7 }, { "b": "bWJfdXNlcg==", "s": "mb_user", "ii": 16, "i": 8 }, { "b": "MTUxNTE=", "s": "15151", "ii": 17, "i": 9 }, { "b": "cmVwbHk=", "s": "reply", "ii": 18, "i": 10 }, { "b": "bnVsbA==", "s": "null", "ii": 19, "i": 11 }, { "b": "dHlwZQ==", "s": "type", "ii": 20, "i": 12 }, { "b": "cG9zdA==", "s": "post", "ii": 21, "i": 13 }, { "b": "dGltZXN0YW1w", "s": "timestamp", "ii": 22, "i": 14 }, { "b": "bnVsbA==", "s": "null", "ii": 23, "i": 15 }, { "b": "YXBw", "s": "app", "ii": 24, "i": 16 }, { "b": "dHdldGNo", "s": "twetch", "ii": 25, "i": 17 }, { "b": "aW52b2ljZQ==", "s": "invoice", "ii": 26, "i": 18 }, { "b": "MWNhY2JmMzMtNjViYi00NTI4LTk0MmMtZjUzODM0YmM5OWJh", "s": "1cacbf33-65bb-4528-942c-f53834bc99ba", "ii": 27, "i": 19 } ], "i": 2 }, { "cell": [ { "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "ii": 29, "i": 0 }, { "b": "QklUQ09JTl9FQ0RTQQ==", "s": "BITCOIN_ECDSA", "ii": 30, "i": 1 }, { "b": "MTN2NlBFRVk1azdNdEFralk0Rld4Nm9XakFYTjh3UFRwZw==", "s": "13v6PEEY5k7MtAkjY4FWx6oWjAXN8wPTpg", "ii": 31, "i": 2 }, { "b": "SUY3UEVDQWd0MTlPNFdQMU43VzlTU1VQaXFjTDNBWE9RR3RsbWgvUjA3b0JYOHpkeUlETFBmRmhDZkFYTkNEcDh2eTVib3F5c3pKUkpNeHlaZ3Fzc1hVPQ==", "s": "IF7PECAgt19O4WP1N7W9SSUPiqcL3AXOQGtlmh/R07oBX8zdyIDLPfFhCfAXNCDp8vy5boqyszJRJMxyZgqssXU=", "ii": 32, "i": 3 } ], "i": 3 } ], "e": { "v": 0, "i": 0, "a": "false" } 	}, 	{ "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "Ron3JFB3Ecc4h7eCRqjoglw5IBM=", "s": "F��$Pw\u0011�8���F��\\9 \u0013", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 968, "i": 1, "a": "17RyaS4LcXxqiupvTSbkgd1kgF42DiA64h" } 	}, 	{ "i": 2, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "BRhv8HEO0AQinmRMBlOymFxkiiM=", "s": "\u0005\u0018o�q\u000e�\u0004\"�dL\u0006S��\\d�#", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 8711, "i": 2, "a": "1Twetcht1cTUxpdDoX5HQRpoXeuupAdyf" } 	}, 	{ "i": 3, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "g83r1UQOPksOnIJfTGI3/MPO17w=", "s": "����D\u000e>K\u000e��_Lb7���׼", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 421805, "i": 3, "a": "1D1vCnf6537YAq4BjqZqMJdtKfFnX5yEAz" } 	} ], "lock": 0, "blk": { 	"i": 633484, 	"h": "000000000000000001cdfe7550e9f77bec9d10bc38e4a840f96db85eb64a787e", 	"t": 1588629096 }, "i": 3453	}	`

	err := bobData.FromString(bobString)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	// TODO - the above data does not contain the actual B information from the tx it has been stripped out
	t.Logf("Stuff %+v", bobData.Out[0].Tape[1])
	bData := New()
	err = bData.FromTape(bobData.Out[0].Tape[1])
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	dataURI := bData.DataURI()
	bitfsURL := BitFsURL(bobData.Tx.H, 0, 3)

	// Build fileName from fullPath
	var fileURL *url.URL
	fileURL, err = url.Parse(bitfsURL)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName := segments[len(segments)-1]

	// Create blank file
	var file *os.File
	file, err = os.Create(fileName)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	// Put content on file
	var req *http.Request
	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, bitfsURL, nil)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	t.Log(base64.StdEncoding.EncodeToString(body))

	var size int64
	size, err = io.Copy(file, resp.Body)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	defer func() {
		_ = file.Close()
	}()

	t.Logf("Just Downloaded a file %s with size %d\n\n", fileName, size)

	t.Log("DataURI", dataURI, bData.Data.UTF8, bitfsURL)
}
