package aip

// Example BOB tx (valid signature and address)
const sampleValidBobTx = `{ "_id": "5f08ddb0f797435fbff1ddf0", "tx": { "h": "744a55a8637aa191aa058630da51803abbeadc2de3d65b4acace1f5f10789c5b" }, 
"out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "i": 0, "ii": 0 }, { "op": 106, "ops": "OP_RETURN", "i": 1, "ii": 1 } ], "i": 0 }, 
{ "cell": [ { "s": "1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT", "h": "31424150537561506e66476e53424d33474c56397968785564596534764762644d54", 
"b": "MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==", "i": 0, "ii": 2 }, { "s": "ATTEST", "h": "415454455354", "b": "QVRURVNU", "i": 1, "ii": 3 },
{ "s": "cf39fc55da24dc23eff1809e6e6cf32a0fe6aecc81296543e9ac84b8c501bac5", 
"h": "63663339666335356461323464633233656666313830396536653663663332613066653661656363383132393635343365396163383462386335303162616335",
"b": "Y2YzOWZjNTVkYTI0ZGMyM2VmZjE4MDllNmU2Y2YzMmEwZmU2YWVjYzgxMjk2NTQzZTlhYzg0YjhjNTAxYmFjNQ==", "i": 2, "ii": 4 },
{ "s": "0", "h": "30", "b": "MA==", "i": 3, "ii": 5 } ], "i": 1 }, { "cell": [ { "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", 
"h": "313550636948473232534e4c514a584d6f5355615756693757537163376843667661", "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "i": 0, "ii": 7 },
{ "s": "BITCOIN_ECDSA", "h": "424954434f494e5f4543445341", "b": "QklUQ09JTl9FQ0RTQQ==", "i": 1, "ii": 8 }, { "s": "134a6TXxzgQ9Az3w8BcvgdZyA5UqRL89da", 
"h": "31333461365458787a675139417a33773842637667645a7941355571524c38396461", "b": "MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==", "i": 2, "ii": 9 },
{ "s": "\u001f�nm�3坨\u001b�{\u001f\t��\u0000��(ӏ��h�D��o\u000b\u0006�$�(\u001a�'i��_�\u0006YA\"\f��ޚ` + "`" + `/U.\u0012�^W�\n", 
"h": "1fe96e6df733e59da81bc07b1f098ff19fad00b3fe28d38f81e768ed44d7c16f0b06932480281ab42769bdbb5fef065941220ccfcdde9a602f552e12dc5e57d70a", 
"b": "H+lubfcz5Z2oG8B7HwmP8Z+tALP+KNOPgedo7UTXwW8LBpMkgCgatCdpvbtf7wZZQSIMz83emmAvVS4S3F5X1wo=", "i": 3, "ii": 10 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, 
{ "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "i": 0, "ii": 0 }, { "op": 169, "ops": "OP_HASH160", "i": 1, "ii": 1 }, 
{ "s": "�\no;L˺��E\t^��{i\u0011}", "h": "d27f0a6f3b4ccbbacaf945095ed3eeb97b69117d", "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "i": 2, "ii": 2 },
{ "op": 136, "ops": "OP_EQUALVERIFY", "i": 3, "ii": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "i": 4, "ii": 4 } ], "i": 0 } ], 
"e": { "v": 14492205, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "timestamp": 1594416560292 }`

// Example BOB tx (Invalid Address)
const sampleInvalidBobTx = `{ "_id": "5f08ddb0f797435fbff1ddf0", "tx": { "h": "744a55a8637aa191aa058630da51803abbeadc2de3d65b4acace1f5f10789c5b" }, 
"out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "i": 0, "ii": 0 }, { "op": 106, "ops": "OP_RETURN", "i": 1, "ii": 1 } ], "i": 0 }, 
{ "cell": [ { "s": "1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT", "h": "31424150537561506e66476e53424d33474c56397968785564596534764762644d54", 
"b": "MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==", "i": 0, "ii": 2 }, { "s": "ATTEST", "h": "415454455354", "b": "QVRURVNU", "i": 1, "ii": 3 },
{ "s": "cf39fc55da24dc23eff1809e6e6cf32a0fe6aecc81296543e9ac84b8c501bac5", 
"h": "63663339666335356461323464633233656666313830396536653663663332613066653661656363383132393635343365396163383462386335303162616335",
"b": "Y2YzOWZjNTVkYTI0ZGMyM2VmZjE4MDllNmU2Y2YzMmEwZmU2YWVjYzgxMjk2NTQzZTlhYzg0YjhjNTAxYmFjNQ==", "i": 2, "ii": 4 },
{ "s": "0", "h": "30", "b": "MA==", "i": 3, "ii": 5 } ], "i": 1 }, { "cell": [ { "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", 
"h": "313550636948473232534e4c514a584d6f5355615756693757537163376843667661", "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "i": 0, "ii": 7 },
{ "s": "BITCOIN_ECDSA", "h": "424954434f494e5f4543445341", "b": "QklUQ09JTl9FQ0RTQQ==", "i": 1, "ii": 8 }, { "s": "invalid-address", 
"h": "31333461365458787a675139417a33773842637667645a7941355571524c38396461", "b": "MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==", "i": 2, "ii": 9 },
{ "s": "\u001f�nm�3坨\u001b�{\u001f\t��\u0000��(ӏ��h�D��o\u000b\u0006�$�(\u001a�'i��_�\u0006YA\"\f��ޚ` + "`" + `/U.\u0012�^W�\n", 
"h": "1fe96e6df733e59da81bc07b1f098ff19fad00b3fe28d38f81e768ed44d7c16f0b06932480281ab42769bdbb5fef065941220ccfcdde9a602f552e12dc5e57d70a", 
"b": "H+lubfcz5Z2oG8B7HwmP8Z+tALP+KNOPgedo7UTXwW8LBpMkgCgatCdpvbtf7wZZQSIMz83emmAvVS4S3F5X1wo=", "i": 3, "ii": 10 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, 
{ "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "i": 0, "ii": 0 }, { "op": 169, "ops": "OP_HASH160", "i": 1, "ii": 1 }, 
{ "s": "�\no;L˺��E\t^��{i\u0011}", "h": "d27f0a6f3b4ccbbacaf945095ed3eeb97b69117d", "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "i": 2, "ii": 2 },
{ "op": 136, "ops": "OP_EQUALVERIFY", "i": 3, "ii": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "i": 4, "ii": 4 } ], "i": 0 } ], 
"e": { "v": 14492205, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "timestamp": 1594416560292 }`

// var sampleBobTx2 = `{ "_id": "5ec76f5560876f1667783bbf", "tx": { "h": "6c6e52da3f16f6a03a9ee5bfd68dd6a9fb7fce16fc66f137a265a4bf7cbb4cba"},"in":[{"i":0,"tape":[{"cell":[{"b":"MEQCIGWhlEXTARNNmAbSrAYUPdTvUAtY47jU9/iJ1mGyPRjjAiAkkhaqdRMmxQY7c4Ci4xbrFPU73F9ZGjrxP8A5FJ9mPEE=","s":"0D\u0002 e��E�\u0001\u0013M�\u0006Ҭ\u0006\u0014=��P\u000bX������a�=\u0018�\u0002 $�\u0016�u\u0013&�\u0006;s���\u0016�\u0014�;�_Y\u001a:�?�9\u0014�f<A","ii":0,"i":0},{"b":"BED/sziEj3i/u3i5tKgsIx3HKM7vQrNBJQyEupnPRYvyrwCV31Rb7z0o5xfNvwEQKhxyXGla3+QHSGGVGFdN8ig=","s":"\u0004@��8��x��x���,#\u001d�(��B�A%\f����E��\u0000��T[�=(�\u0017Ϳ\u0001\u0010*\u001cr\\iZ��\u0007Ha�\u0018WM�(","ii":1,"i":1}],"i":0}],"e":{"h":"9499b298b2d6e85b460530847b09b585875fdf2c1ab76b0175499906ae208d24","i":1,"a":"1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk"},"seq":4294967295}],"out":[{"i":0,"tape":[{"cell":[{"op":0,"ops":"OP_0","ii":0,"i":0},{"op":106,"ops":"OP_RETURN","ii":1,"i":1}],"i":0},{"cell":[{"b":"MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==","s":"1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT","ii":2,"i":0},{"b":"QVRURVNU","s":"ATTEST","ii":3,"i":1},{"b":"OWI3ZDViOTBjMmFjYTU5OGYyOTkwYmIwNmRjMmU1ZGZkNmRiMjFjMTM4ZDk2YjNhMzJkYmEyNWQ0Zjc1ZWYxYw==","s":"9b7d5b90c2aca598f2990bb06dc2e5dfd6db21c138d96b3a32dba25d4f75ef1c","ii":4,"i":2},{"b":"MA==","s":"0","ii":5,"i":3}],"i":1},{"cell":[{"b":"MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==","s":"15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva","ii":7,"i":0},{"b":"QklUQ09JTl9FQ0RTQQ==","s":"BITCOIN_ECDSA","ii":8,"i":1},{"b":"MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==","s":"134a6TXxzgQ9Az3w8BcvgdZyA5UqRL89da","ii":9,"i":2},{"b":"IMPuiSTsxdaDP38JxEpsoSCr3hwosWBYpzMGeabXRogDJMEIHNiB2x4G8kxb4PYr+Fbf0vGUsPQLa7DBOViE+Nc=","s":" ��$��փ?\t�Jl� ��\u001c(�` + "`" + `X�3\u0006y��F�\u0003$�\b\u001c؁�\u001e\u0006�L[��+�V����\u000bk��9X���","ii":10,"i":3}],"i":2}],"e":{"v":0,"i":0,"a":"false"}},{"i":1,"tape":[{"cell":[{"op":118,"ops":"OP_DUP","ii":0,"i":0},{"op":169,"ops":"OP_HASH160","ii":1,"i":1},{"b":"0n8KbztMy7rK+UUJXtPuuXtpEX0=","s":"�\no;L˺��E\t^��{i\u0011}","ii":2,"i":2},{"op":136,"ops":"OP_EQUALVERIFY","ii":3,"i":3},{"op":172,"ops":"OP_CHECKSIG","ii":4,"i":4}],"i":0}],"e":{"v":15830368,"i":1,"a":"1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk"}}],"lock":0,"blk":{"i":628688,"h":"00000000000000000264d87d624a8d9758c37767d5cc0ca3650df7b5aa94012d","t":1585698239},"i":3244}`
// const sampleBobTx = `{ "_id": "5ece87335fdea4165b05796f", "tx": { "h": "232e9b089b95c871c62e87a179e4cbfc8dad6c864e39db8d38aaf9ef1a7f44e0" }, "in": [ { "i": 0, "tape": [ { "cell": [ { "b": "MEQCIH2pVbX6hGoxGt5Fa+ocNedSvdxhnDTwut0S29xFdLpUAiALY63PXcXptMuzB45kwd9Nn5XMVuaPXlBmUe8sGMll7kE=", "s": "0D\u0002 }�U���j1\u001a�Ek�\u001c5�R��a�4��\u0012��Et�T\u0002 \u000bc��]��˳\u0007�d��M���V�^PfQ�,\u0018�e�A", "ii": 0, "i": 0 }, { "b": "BED/sziEj3i/u3i5tKgsIx3HKM7vQrNBJQyEupnPRYvyrwCV31Rb7z0o5xfNvwEQKhxyXGla3+QHSGGVGFdN8ig=", "s": "\u0004@��8��x��x���,#\u001d�(��B�A%\f����E��\u0000��T[�=(�\u0017Ϳ\u0001\u0010*\u001cr\\iZ��\u0007Ha�\u0018WM�(", "ii": 1, "i": 1 } ], "i": 0 } ], "e": { "h": "e9227531f5468715ea9aaf2fec2c95ec8e20af31c64acf81e6463c7eebfecbbc", "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" }, "seq": 4294967295 } ], "out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "ii": 0, "i": 0 }, { "op": 106, "ops": "OP_RETURN", "ii": 1, "i": 1 } ], "i": 0 }, { "cell": [ { "b": "MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==", "s": "1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT", "ii": 2, "i": 0 }, { "b": "QVRURVNU", "s": "ATTEST", "ii": 3, "i": 1 }, { "b": "MWI3Y2VmZGY1ZjBlZTc0Mjk3ZjUyZGQ3ZTAzYTA2ZTg1YjUyZWEzOWJlNmQ4Zjg5ZTU3YzRmYjgyNzRlNzU2Nw==", "s": "1b7cefdf5f0ee74297f52dd7e03a06e85b52ea39be6d8f89e57c4fb8274e7567", "ii": 4, "i": 2 }, { "b": "MA==", "s": "0", "ii": 5, "i": 3 } ], "i": 1 }, { "cell": [ { "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "ii": 7, "i": 0 }, { "b": "QklUQ09JTl9FQ0RTQQ==", "s": "BITCOIN_ECDSA", "ii": 8, "i": 1 }, { "b": "MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==", "s": "134a6TXxzgQ9Az3w8BcvgdZyA5UqRL89da", "ii": 9, "i": 2 }, { "b": "IDdIMY7qZa6nZeDX5pcDTpsvmlkAECMjgoa7XROnoqmpMuciBUWO7bMBugapIkpjLy5O/WxXfbKzyvEqn2Y/OAQ=", "s": " 7H1��e��e���\u0003N�/�Y\u0000\u0010##���]\u0013����2�\"\u0005E���\u0001�\u0006�\"Jc/.N�lW}����*�f?8\u0004", "ii": 10, "i": 3 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, { "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "s": "�\no;L˺��E\t^��{i\u0011}", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 14498894, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "blk": { "i": 634446, "h": "000000000000000004d17661eb7c48c65103191fb5f7e4b94c363cd86acc0741", "t": 1589217048 }, "i": 89 }`
// const sampleBobTx = `{ "_id": "5ec7700660876f166778419a", "tx": { "h": "447bdbf0d910fb200ec604dde3d3745ccb285b8b8967c624f273dc8fc9bd2491" }, "in": [ { "i": 0, "tape": [ { "cell": [ { "b": "MEQCIHLammTixe8nWhDKF/0SONGMaRWCBg9/azVUiJlSMLW4AiBMJ25zptdWqwGNKkeoxdHCZJkwDdAqGJN+1ZOx/yn3tUE=", "s": "0D\u0002 rښd���'Z\u0010�\u0017�\u00128ьi\u0015�\u0006\u000fk5T��R0��\u0002 L'ns��V�\u0001�*G����d�0\r�*\u0018�~Փ��)��A", "ii": 0, "i": 0 }, { "b": "BED/sziEj3i/u3i5tKgsIx3HKM7vQrNBJQyEupnPRYvyrwCV31Rb7z0o5xfNvwEQKhxyXGla3+QHSGGVGFdN8ig=", "s": "\u0004@��8��x��x���,#\u001d�(��B�A%\f����E��\u0000��T[�=(�\u0017Ϳ\u0001\u0010*\u001cr\\iZ��\u0007Ha�\u0018WM�(", "ii": 1, "i": 1 } ], "i": 0 } ], "e": { "h": "15e89012d156be10d2ba682d0f4a5641aa91b7c5804218652fe5c61b40f061bb", "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" }, "seq": 4294967295 } ], "out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "ii": 0, "i": 0 }, { "op": 106, "ops": "OP_RETURN", "ii": 1, "i": 1 } ], "i": 0 }, { "cell": [ { "b": "MTlIeGlnVjRReUJ2M3RIcFFWY1VFUXlxMXB6WlZkb0F1dA==", "s": "19HxigV4QyBv3tHpQVcUEQyq1pzZVdoAut", "ii": 2, "i": 0 }, { "b": "eyJpdiI6IjI5Nzg5NzAwODkwNjQxN2Q5NTczIiwic2FsdCI6IjgwYjdlZjNiZWEwYTRmZjYxMzQxIiwia2VtdGFnIjoiWHdyZU1vZHRjejR5OGlFbDhZQ0dzRjRYSHVjRmVsTlJsNFF4a0x3RkJvdHRXenZmY2dPVVQ2enNHalJNdFFTb1FsS1pzUUt4U0RGYUxLZkdscG9saWc9PSIsImN0IjoiZlR6M1dobEFFemZXWldKZno2azl1cmVZdFBHQW5PR1V0NEV3TjhYZ1JRNG9Sb1RGTm9oMXZwMWxJWjBUS2w4SFNnQU5qNk1iak1JVk1XajdPSU09In0=", "s": "{\"iv\":\"297897008906417d9573\",\"salt\":\"80b7ef3bea0a4ff61341\",\"kemtag\":\"XwreModtcz4y8iEl8YCGsF4XHucFelNRl4QxkLwFBottWzvfcgOUT6zsGjRMtQSoQlKZsQKxSDFaLKfGlpolig==\",\"ct\":\"fTz3WhlAEzfWZWJfz6k9ureYtPGAnOGUt4EwN8XgRQ4oRoTFNoh1vp1lIZ0TKl8HSgANj6MbjMIVMWj7OIM=\"}", "ii": 3, "i": 1 }, { "b": "YXBwbGljYXRpb24vYWVzLWVuY3J5cHRlZDtjb250ZW50LXR5cGU9InRleHQvaHRtbCI=", "s": "application/aes-encrypted;content-type=\"text/html\"", "ii": 4, "i": 2 }, { "b": "dXRmLTg=", "s": "utf-8", "ii": 5, "i": 3 }, { "b": "AA==", "s": "\u0000", "ii": 6, "i": 4 } ], "i": 1 }, { "cell": [ { "b": "MUxDUGt1MlJ6TmlBeDZtV2YzZk5hZWZid3VEQ3JWN21ZSg==", "s": "1LCPku2RzNiAx6mWf3fNaefbwuDCrV7mYJ", "ii": 8, "i": 0 }, { "ii": 9, "i": 1, "ls": "{\"iv\":\"aca5b566aa053087c3cd\",\"salt\":\"51d774ba70238ca75270\",\"kemtag\":\"laTadPWcI9sQWMdlsHGQhKhzY1/p+zscE3+0Cw4CUscKVewpgb7+89qUCWpXuP2M2/XgwtAxXueYIp84dYQydg==\",\"ct\":\"R2RUkRBws3sZKpTC+xMGzUCHJRO8+uBStTQm1MDKj2Wmbq7F6w9bXrj3sQAcTo9v1rRmkivLRRPdVUp6f+sb0AOmapdb60fXXKFvumc+12QXUb5alfiWWTeb9u830oeFfAUj0Z7CmRzHAEDDK4WeW2nkmCdfBpG/YnOWzqXya57pwKEAc9zBEWAUQI1Qfz9zDbGP06SgyKx1ALJzYKY/zQSi7ioU2luY/f6UpVyDt7+j/201JDmIPHzgHO0P3RKCmQ/94c/3+nOdILr4jTP0oXwvuMMjA1L2i+KMeSJfJbtG5CYgq6s5DvhMA1NkWiYrxp5L37YvpMbMdwLN2fyBDHuiPukvwa7xrDra00mgULHCA0/c7Uy+UzZZhBq9GJqRdToAhjrvhoqlNb5ctq16SOYJzSG9agOlRCUHrsnJUZdk/rhX0NM=\"}", "lb": "eyJpdiI6ImFjYTViNTY2YWEwNTMwODdjM2NkIiwic2FsdCI6IjUxZDc3NGJhNzAyMzhjYTc1MjcwIiwia2VtdGFnIjoibGFUYWRQV2NJOXNRV01kbHNIR1FoS2h6WTEvcCt6c2NFMyswQ3c0Q1VzY0tWZXdwZ2I3Kzg5cVVDV3BYdVAyTTIvWGd3dEF4WHVlWUlwODRkWVF5ZGc9PSIsImN0IjoiUjJSVWtSQndzM3NaS3BUQyt4TUd6VUNISlJPOCt1QlN0VFFtMU1ES2oyV21icTdGNnc5YlhyajNzUUFjVG85djFyUm1raXZMUlJQZFZVcDZmK3NiMEFPbWFwZGI2MGZYWEtGdnVtYysxMlFYVWI1YWxmaVdXVGViOXU4MzBvZUZmQVVqMFo3Q21SekhBRURESzRXZVcybmttQ2RmQnBHL1luT1d6cVh5YTU3cHdLRUFjOXpCRVdBVVFJMVFmejl6RGJHUDA2U2d5S3gxQUxKellLWS96UVNpN2lvVTJsdVkvZjZVcFZ5RHQ3K2ovMjAxSkRtSVBIemdITzBQM1JLQ21RLzk0Yy8zK25PZElMcjRqVFAwb1h3dnVNTWpBMUwyaStLTWVTSmZKYnRHNUNZZ3E2czVEdmhNQTFOa1dpWXJ4cDVMMzdZdnBNYk1kd0xOMmZ5QkRIdWlQdWt2d2E3eHJEcmEwMG1nVUxIQ0EwL2M3VXkrVXpaWmhCcTlHSnFSZFRvQWhqcnZob3FsTmI1Y3RxMTZTT1lKelNHOWFnT2xSQ1VIcnNuSlVaZGsvcmhYME5NPSJ9" } ], "i": 2 }, { "cell": [ { "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "ii": 11, "i": 0 }, { "b": "QklUQ09JTl9FQ0RTQQ==", "s": "BITCOIN_ECDSA", "ii": 12, "i": 1 }, { "b": "MUQyUlJOazE4bnh0WkFmN29OcDhqU1F4WWZmQWVHTk54Wg==", "s": "1D2RRNk18nxtZAf7oNp8jSQxYffAeGNNxZ", "ii": 13, "i": 2 }, { "b": "IMRNZ8W1qTBKrtoao5oGTtlNUz7ZWbx2gmjDUFgJrcxZOmehuukAIkpZB+4NrHyIP5nd4lUwgt+atR6F3nWhLWc=", "s": " �Mgŵ�0J��\u001a��\u0006N�MS>�Y�v�h�PX\t��Y:g���\u0000\"JY\u0007�\r�|�?���U0�ߚ�\u001e��u�-g", "ii": 14, "i": 3 }, { "b": "AQ==", "s": "\u0001", "ii": 15, "i": 4 } ], "i": 3 } ], "e": { "v": 0, "i": 0, "a": "false" } }, { "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "s": "�\no;L˺��E\t^��{i\u0011}", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 15827826, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "blk": { "i": 628691, "h": "000000000000000002d6a818cc26adceedeace1829fa5c2d5277f1fe81c6ce90", "t": 1585702224 }, "i": 59 }`
// const sampleBobTx = `{ "_id": "5f08ddeed1352a2c3432f4db", "tx": { "h": "26b754e6fdf04121b8d91160a0b252a22ae30204fc552605b7f6d3f08419f29e" }, "in": [ { "i": 0, "seq": 4294967295, "tape": [ { "cell": [ { "s": "0E\u0002!\u0000����;�Z��\b\th�&���5����6��` + "`" + `\u0016�Z�N\u0002 WUI\u001bz)\nE{\u001f��0�g�꨻*}\u0018QV��dO�D@�A", "h": "3045022100afbbffff3bb55aaec20809689026acbccf35bcb4e2f29c36aaf86016d85abe4e02205755491b7a290a457b1fbea2308567ddeaa8bb2a7d185156a1f3644f854440d941", "b": "MEUCIQCvu///O7VarsIICWiQJqy8zzW8tOLynDaq+GAW2Fq+TgIgV1VJG3opCkV7H76iMIVn3eqouyp9GFFWofNkT4VEQNlB", "i": 0, "ii": 0 }, { "s": "\u0004@��8��x��x���,#\u001d�(��B�A%\f����E��\u0000��T[�=(�\u0017Ϳ\u0001\u0010*\u001cr\\iZ��\u0007Ha�\u0018WM�(", "h": "0440ffb338848f78bfbb78b9b4a82c231dc728ceef42b341250c84ba99cf458bf2af0095df545bef3d28e717cdbf01102a1c725c695adfe40748619518574df228", "b": "BED/sziEj3i/u3i5tKgsIx3HKM7vQrNBJQyEupnPRYvyrwCV31Rb7z0o5xfNvwEQKhxyXGla3+QHSGGVGFdN8ig=", "i": 1, "ii": 1 } ], "i": 0 } ], "e": { "h": "744a55a8637aa191aa058630da51803abbeadc2de3d65b4acace1f5f10789c5b", "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "i": 0, "ii": 0 }, { "op": 106, "ops": "OP_RETURN", "i": 1, "ii": 1 } ], "i": 0 }, { "cell": [ { "s": "1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT", "h": "31424150537561506e66476e53424d33474c56397968785564596534764762644d54", "b": "MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==", "i": 0, "ii": 2 }, { "s": "ATTEST", "h": "415454455354", "b": "QVRURVNU", "i": 1, "ii": 3 }, { "s": "16ca90ce3c6347132adba40aa0d5faa3b2bf2015678ffc63db1511b676885e25", "h": "31366361393063653363363334373133326164626134306161306435666161336232626632303135363738666663363364623135313162363736383835653235", "b": "MTZjYTkwY2UzYzYzNDcxMzJhZGJhNDBhYTBkNWZhYTNiMmJmMjAxNTY3OGZmYzYzZGIxNTExYjY3Njg4NWUyNQ==", "i": 2, "ii": 4 }, { "s": "0", "h": "30", "b": "MA==", "i": 3, "ii": 5 } ], "i": 1 }, { "cell": [ { "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "h": "313550636948473232534e4c514a584d6f5355615756693757537163376843667661", "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "i": 0, "ii": 7 }, { "s": "BITCOIN_ECDSA", "h": "424954434f494e5f4543445341", "b": "QklUQ09JTl9FQ0RTQQ==", "i": 1, "ii": 8 }, { "s": "134a6TXxzgQ9Az3w8BcvgdZyA5UqRL89da", "h": "31333461365458787a675139417a33773842637667645a7941355571524c38396461", "b": "MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==", "i": 2, "ii": 9 }, { "s": "\u001f�V���j{k�\u0010ҕ�QA�]�Ӛ` + "`" + `7N����^���)YΓ\u001f@�qWcH}�V��Y�\u0019F�C�V�@�\r�a�", "h": "1fc756c3fcc76a7b6bcf10d295a75141ef5dbbd39a60374ea796eb92d85e84a0a32959ce931f40dc715763487de7a856acca59fc19468343b4569340d20d9761ed", "b": "H8dWw/zHantrzxDSladRQe9du9OaYDdOp5brkthehKCjKVnOkx9A3HFXY0h956hWrMpZ/BlGg0O0VpNA0g2XYe0=", "i": 3, "ii": 10 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, { "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "i": 0, "ii": 0 }, { "op": 169, "ops": "OP_HASH160", "i": 1, "ii": 1 }, { "s": "�\no;L˺��E\t^��{i\u0011}", "h": "d27f0a6f3b4ccbbacaf945095ed3eeb97b69117d", "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "i": 2, "ii": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "i": 3, "ii": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "i": 4, "ii": 4 } ], "i": 0 } ], "e": { "v": 14491552, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "timestamp": 1594416622135 }`
// const sampleBobTx = `{ "_id": "5f08ddeed1352a2c3432f4db", "tx": { "h": "26b754e6fdf04121b8d91160a0b252a22ae30204fc552605b7f6d3f08419f29e" }, "in": [ { "i": 0, "seq": 4294967295, "tape": [ { "cell": [ { "s": "0E\u0002!\u0000����;�Z��\b\th�&���5����6��` + "`" + `\u0016�Z�N\u0002 WUI\u001bz)\nE{\u001f��0�g�꨻*}\u0018QV��dO�D@�A", "h": "3045022100afbbffff3bb55aaec20809689026acbccf35bcb4e2f29c36aaf86016d85abe4e02205755491b7a290a457b1fbea2308567ddeaa8bb2a7d185156a1f3644f854440d941", "b": "MEUCIQCvu///O7VarsIICWiQJqy8zzW8tOLynDaq+GAW2Fq+TgIgV1VJG3opCkV7H76iMIVn3eqouyp9GFFWofNkT4VEQNlB", "i": 0, "ii": 0 }, { "s": "\u0004@��8��x��x���,#\u001d�(��B�A%\f����E��\u0000��T[�=(�\u0017Ϳ\u0001\u0010*\u001cr\\iZ��\u0007Ha�\u0018WM�(", "h": "0440ffb338848f78bfbb78b9b4a82c231dc728ceef42b341250c84ba99cf458bf2af0095df545bef3d28e717cdbf01102a1c725c695adfe40748619518574df228", "b": "BED/sziEj3i/u3i5tKgsIx3HKM7vQrNBJQyEupnPRYvyrwCV31Rb7z0o5xfNvwEQKhxyXGla3+QHSGGVGFdN8ig=", "i": 1, "ii": 1 } ], "i": 0 } ], "e": { "h": "744a55a8637aa191aa058630da51803abbeadc2de3d65b4acace1f5f10789c5b", "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "i": 0, "ii": 0 }, { "op": 106, "ops": "OP_RETURN", "i": 1, "ii": 1 } ], "i": 0 }, { "cell": [ { "s": "1BAPSuaPnfGnSBM3GLV9yhxUdYe4vGbdMT", "h": "31424150537561506e66476e53424d33474c56397968785564596534764762644d54", "b": "MUJBUFN1YVBuZkduU0JNM0dMVjl5aHhVZFllNHZHYmRNVA==", "i": 0, "ii": 2 }, { "s": "ATTEST", "h": "415454455354", "b": "QVRURVNU", "i": 1, "ii": 3 }, { "s": "16ca90ce3c6347132adba40aa0d5faa3b2bf2015678ffc63db1511b676885e25", "h": "31366361393063653363363334373133326164626134306161306435666161336232626632303135363738666663363364623135313162363736383835653235", "b": "MTZjYTkwY2UzYzYzNDcxMzJhZGJhNDBhYTBkNWZhYTNiMmJmMjAxNTY3OGZmYzYzZGIxNTExYjY3Njg4NWUyNQ==", "i": 2, "ii": 4 }, { "s": "0", "h": "30", "b": "MA==", "i": 3, "ii": 5 } ], "i": 1 }, { "cell": [ { "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "h": "313550636948473232534e4c514a584d6f5355615756693757537163376843667661", "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "i": 0, "ii": 7 }, { "s": "BITCOIN_ECDSA", "h": "424954434f494e5f4543445341", "b": "QklUQ09JTl9FQ0RTQQ==", "i": 1, "ii": 8 }, { "s": "134a6TXxzgQ9Az3w8BcvgdZyA5UqRL89da", "h": "31333461365458787a675139417a33773842637667645a7941355571524c38396461", "b": "MTM0YTZUWHh6Z1E5QXozdzhCY3ZnZFp5QTVVcVJMODlkYQ==", "i": 2, "ii": 9 }, { "s": "\u001f�V���j{k�\u0010ҕ�QA�]�Ӛ` + "`" + `7N����^���)YΓ\u001f@�qWcH}�V��Y�\u0019F�C�V�@�\r�a�", "h": "1fc756c3fcc76a7b6bcf10d295a75141ef5dbbd39a60374ea796eb92d85e84a0a32959ce931f40dc715763487de7a856acca59fc19468343b4569340d20d9761ed", "b": "H8dWw/zHantrzxDSladRQe9du9OaYDdOp5brkthehKCjKVnOkx9A3HFXY0h956hWrMpZ/BlGg0O0VpNA0g2XYe0=", "i": 3, "ii": 10 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, { "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "i": 0, "ii": 0 }, { "op": 169, "ops": "OP_HASH160", "i": 1, "ii": 1 }, { "s": "�\no;L˺��E\t^��{i\u0011}", "h": "d27f0a6f3b4ccbbacaf945095ed3eeb97b69117d", "b": "0n8KbztMy7rK+UUJXtPuuXtpEX0=", "i": 2, "ii": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "i": 3, "ii": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "i": 4, "ii": 4 } ], "i": 0 } ], "e": { "v": 14491552, "i": 1, "a": "1LC16EQVsqVYGeYTCrjvNf8j28zr4DwBuk" } } ], "lock": 0, "timestamp": 1594416622135 }`
// const sampleBobTx = `{ "_id": "5ed07f4b57cd6b1658b817f7", "tx": { "h": "375e67f427d04e1e1a202be6f27ec33a382d3a655af539c079a9f595ec606bef" }, "in": [ { "i": 0, "tape": [ { "cell": [ { "b": "MEQCIAIKjpbGASg3rUKJwqUPW08rlcf+inWtoaTa6fnDV/gMAiBXl1x2YSZpvLi6OVot1+G23BQbIIViDv09YbFXZy+mBUE=", "s": "0D\u0002 \u0002\n���\u0001(7�B�¥\u000f[O+����u�������W�\f\u0002 W�\\va&i���9Z-���\u0014\u001b �b\u000e�=a�Wg/�\u0005A", "ii": 0, "i": 0 }, { "b": "A1B3SiT3OTW6r9F1cT651UwXkAG64vhQJeSnsCrL9fA4", "s": "\u0003PwJ$�95���uq>��L\u0017�\u0001���P%䧰*���8", "ii": 1, "i": 1 } ], "i": 0 } ], "e": { "h": "48e93234cb6aaf1098c4195164e426c67b0104b744758f146e0d1496bb7d6ebf", "i": 5, "a": "1Bpx4FdsENLcFgvkpEmBVu1o2AgqW2Ye5j" }, "seq": 4294967295 } ], "out": [ { "i": 0, "tape": [ { "cell": [ { "op": 0, "ops": "OP_0", "ii": 0, "i": 0 }, { "op": 106, "ops": "OP_RETURN", "ii": 1, "i": 1 } ], "i": 0 }, { "cell": [ { "b": "MUxvdmVGN3FRaWpwamFzY1B5dEhvcjJ1U0VFakhISDhZQg==", "s": "1LoveF7qQijpjascPytHor2uSEEjHHH8YB", "ii": 2, "i": 0 }, { "b": "NWYwNDcwMTEwZTExNTIwNzlmMjU2MjNhZDJjYzNhYmRmOGU0ODU4MzFkMGI1MzJhYzkxNWY3Zjc0MGQ5NWFiMQ==", "s": "5f0470110e1152079f25623ad2cc3abdf8e485831d0b532ac915f7f740d95ab1", "ii": 3, "i": 1 }, { "b": "dHdldGNo", "s": "twetch", "ii": 4, "i": 2 }, { "b": "YmUxMzI0NzYtM2IxZS00Mzk0LTk1YTItODM2Y2I2ZjE3M2I4", "s": "be132476-3b1e-4394-95a2-836cb6f173b8", "ii": 5, "i": 3 } ], "i": 1 }, { "cell": [ { "b": "MTVQY2lIRzIyU05MUUpYTW9TVWFXVmk3V1NxYzdoQ2Z2YQ==", "s": "15PciHG22SNLQJXMoSUaWVi7WSqc7hCfva", "ii": 7, "i": 0 }, { "b": "QklUQ09JTl9FQ0RTQQ==", "s": "BITCOIN_ECDSA", "ii": 8, "i": 1 }, { "b": "MTQ4c3hhY1BYYXBRTmhGQlJuWTlQUFk1cWN5b2lhY0dlcQ==", "s": "148sxacPXapQNhFBRnY9PPY5qcyoiacGeq", "ii": 9, "i": 2 }, { "b": "SVAxUTE0UGxZL1lrRnQwZy9Dd29sdFRKTGhMa2trb3lsOWtJaWhyaE4zWFlQSFRJaDFxbUNteElkTXZ6OTIvaDBnYWdCRnU2ZWtsbUEvMWpaL21DSUV3PQ==", "s": "IP1Q14PlY/YkFt0g/CwoltTJLhLkkkoyl9kIihrhN3XYPHTIh1qmCmxIdMvz92/h0gagBFu6eklmA/1jZ/mCIEw=", "ii": 10, "i": 3 } ], "i": 2 } ], "e": { "v": 0, "i": 0, "a": "false" } }, { "i": 1, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "tyEUOsrxZFF9FqST84M12N7gJXk=", "s": "�!\u0014:��dQ}\u0016���5���%y", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 546, "i": 1, "a": "1HhJHUbJskwHEmxnHE62hVZrbWKAZyXegm" } }, { "i": 2, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "BRhv8HEO0AQinmRMBlOymFxkiiM=", "s": "\u0005\u0018o�q\u000e�\u0004\"�dL\u0006S��\\d�#", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 4718, "i": 2, "a": "1Twetcht1cTUxpdDoX5HQRpoXeuupAdyf" } }, { "i": 3, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "ge8V46sxXUTHoX/MDhoMTUDe5aU=", "s": "��\u0015�1]Dǡ�\u000e\u001a\fM@��", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 6990, "i": 3, "a": "1Cr2ahVvRz5gdNVbKMrNUbqudzeXgVKrx2" } }, { "i": 4, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "Z3WtDwnQdIHbl2gZcRtaVBF/Xng=", "s": "gu�\u000f\t�t�ۗh\u0019q\u001bZT\u0011^x", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 13980, "i": 4, "a": "1AS3a2ocVtMEBFxYiyWKywsZYxRyxxWQCZ" } }, { "i": 5, "tape": [ { "cell": [ { "op": 118, "ops": "OP_DUP", "ii": 0, "i": 0 }, { "op": 169, "ops": "OP_HASH160", "ii": 1, "i": 1 }, { "b": "bnGmFJIqldGih1Hpvw+zKO9pbkQ=", "s": "nq�\u0014�*�Ѣ�Q�\u000f�(�inD", "ii": 2, "i": 2 }, { "op": 136, "ops": "OP_EQUALVERIFY", "ii": 3, "i": 3 }, { "op": 172, "ops": "OP_CHECKSIG", "ii": 4, "i": 4 } ], "i": 0 } ], "e": { "v": 1699347, "i": 5, "a": "1B4yUZC7NkikrXPXEyPEQ6WhVBvYz3asN1" } } ], "lock": 0, "blk": { "i": 635130, "h": "000000000000000002f50329db73a826c97deb0e642e142dc81cdfda4b4a39ba", "t": 1589606545 }, "i": 2690 }`
