# True Go UMDF Prototype

Visual C++ に頼らず CGO だけで Windows ドライバーを書けないか試みる

注意: 動作しないどころか、導入すると頻繁にWindowsがクラッシュしたりハングアップします。安易に試さないようにしてください

# How to build

`"C:\Program Files (x86)\Windows Kits\10\Lib\wdf\umdf\x64\2.15\WdfDriverStubUm.lib"` を `golangs/thirdparty` に複製する。このファイルは再配布不可

```
go build -buildmode=c-shared -o goumdf.dll -ldflags="-v -extldflags '-Wl,--export-all-symbols,-Lthirdparty,-lucrt,-lntdll,-lWdfDriverStubUm,-lntdll -Xlinker --exclude-symbols=_guard_rf_checks_enforced,_guard_icall_checks_enforced,__castguard_slow_path_check_os_handled,__castguard_slow_path_check_nop,__castguard_slow_path_check_fastfail,__castguard_slow_path_check_debugbreak,__castguard_check_failure_os_handled,__castguard_check_failure_nop,__castguard_check_failure_fastfail,__castguard_check_failure_debugbreak,ReadNoFence64,ReadPointerNoFence,_guard_check_icall_nop -Xlinker --script=(このディレクトリのフルパスをエスケープしたもの)\\\\golangs\\\\script.ld
```

`cp golangs/goumdf.dll TrueGoUmdf.dll`

事前に Visual Studio などを使い署名を行い、そのsha1を控える

```
signtool.exe sign /ph /fd "sha256" /sha1 "(当該sha1)" .\TrueGoUmdf.dll
stampinf -d "*" -a "amd64" -u "2.15.0" -v "*" -f .\TrueGoUmdf.inf
inf2cat.exe /os:10_x64 /driver:C:(このディレクトリへのフルパス)\ /uselocaltime
signtool.exe sign /ph /fd "sha256" /sha1 "(当該sha1)" .\truegoumdf.cat
```

ターゲットマシンの `C:\DriverTest\Drivers\` に `TrueGoUmdf.{cat,cer,inf,dll}` をコピーし、証明書をターゲットマシンに読み込ませる

ターゲットマシンで以下を実行するとWindowsがクラッシュするので、事前にいくつか設定しておく
```
C:\DriverTest\devcon install C:\DriverTest\Drivers\TrueGoUmdf.inf ROOT\TrueGoUmdf
```

Target Machine で以下を設定
- WdfVerifier.exe (Windows Driver Kit + Windows SDK) -> Driverの読み込み時に20秒以上待機するように
- Application Verifier (Windows SDK) から `WUDFHost.exe` を有効化
- `devcon install` したときに WinDbg Preview (Microsoft Store) で `WUDFHost.exe` にアタッチ

適宜 `g` で WinDbg Preview をすすめていく。すると `FxDriverEntryUmWorker` で callできず停止し無限ループに陥る