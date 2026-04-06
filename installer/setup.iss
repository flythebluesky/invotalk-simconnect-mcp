; InvoTalk SimConnect MCP Server — Inno Setup Script
; Version is passed from CI via /DMyAppVersion=vX.Y.Z

#ifndef MyAppVersion
  #define MyAppVersion "dev"
#endif

[Setup]
AppName=InvoTalk SimConnect MCP
AppVersion={#MyAppVersion}
AppPublisher=InvoTalk
AppPublisherURL=https://github.com/flythebluesky/invotalk-simconnect-mcp
DefaultDirName={autopf}\InvoTalk SimConnect MCP
DefaultGroupName=InvoTalk SimConnect MCP
OutputBaseFilename=invotalk-simconnect-mcp-setup
Compression=lzma2
SolidCompression=yes
UninstallDisplayIcon={app}\invotalk-simconnect-mcp.exe
ChangesEnvironment=yes
ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible

[Files]
Source: "..\invotalk-simconnect-mcp.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\InvoTalk SimConnect MCP"; Filename: "{app}\invotalk-simconnect-mcp.exe"
Name: "{group}\Uninstall InvoTalk SimConnect MCP"; Filename: "{uninstallexe}"

[Tasks]
Name: "addtopath"; Description: "Add to user PATH"; Flags: unchecked

[Registry]
Root: HKCU; Subkey: "Environment"; ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; Tasks: addtopath; Check: NeedsAddPath(ExpandConstant('{app}'))

[Code]
// NeedsAddPath checks whether the directory is already on PATH.
function NeedsAddPath(Dir: string): Boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', OrigPath) then
  begin
    Result := True;
    exit;
  end;
  Result := Pos(';' + Uppercase(Dir) + ';', ';' + Uppercase(OrigPath) + ';') = 0;
end;

// FindSimConnectDLL checks well-known locations for SimConnect.dll.
// Mirrors the search order in internal/simconnect/dllpath_windows.go — keep in sync.
function FindSimConnectDLL(): string;
var
  Candidate: string;
  SdkPath: string;
begin
  Result := '';

  // Next to installed exe
  Candidate := ExpandConstant('{app}') + '\SimConnect.dll';
  if FileExists(Candidate) then begin Result := Candidate; exit; end;

  // MSFS 2024 SDK env var
  SdkPath := GetEnv('MSFS2024_SDK');
  if SdkPath <> '' then begin
    Candidate := SdkPath + '\SimConnect SDK\lib\SimConnect.dll';
    if FileExists(Candidate) then begin Result := Candidate; exit; end;
  end;

  // MSFS 2020 SDK env var
  SdkPath := GetEnv('MSFS_SDK');
  if SdkPath <> '' then begin
    Candidate := SdkPath + '\SimConnect SDK\lib\SimConnect.dll';
    if FileExists(Candidate) then begin Result := Candidate; exit; end;
  end;

  // Default SDK install paths
  Candidate := 'C:\MSFS 2024 SDK\SimConnect SDK\lib\SimConnect.dll';
  if FileExists(Candidate) then begin Result := Candidate; exit; end;

  Candidate := 'C:\MSFS SDK\SimConnect SDK\lib\SimConnect.dll';
  if FileExists(Candidate) then begin Result := Candidate; exit; end;
end;

procedure CurStepChanged(CurStep: TSetupStep);
var
  DllPath: string;
begin
  if CurStep = ssPostInstall then begin
    DllPath := FindSimConnectDLL();
    if DllPath <> '' then
      MsgBox('SimConnect.dll found at:' + #13#10 + DllPath + #13#10#13#10 + 'The server is ready to connect to MSFS.', mbInformation, MB_OK)
    else
      MsgBox('SimConnect.dll was not found.' + #13#10#13#10 +
        'The server installed successfully, but it needs SimConnect.dll to communicate with Flight Simulator.' + #13#10#13#10 +
        'To install it:' + #13#10 +
        '1. Open MSFS 2024' + #13#10 +
        '2. Go to Options > General > Developers' + #13#10 +
        '3. Enable Developer Mode' + #13#10 +
        '4. Restart the sim, then install the SDK from the Developer menu', mbInformation, MB_OK);
  end;
end;
