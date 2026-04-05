package simconnect

// VarCatalog is a comprehensive list of commonly used SimConnect simulation variables.
var VarCatalog = []VarInfo{
	// Aircraft Info
	{Name: "TITLE", Unit: "string", Category: "Aircraft", Description: "Aircraft title (e.g. Boeing 737-800)", ReadOnly: true},
	{Name: "ATC TYPE", Unit: "string", Category: "Aircraft", Description: "Aircraft ICAO type designator (e.g. B738)", ReadOnly: true},
	{Name: "ATC MODEL", Unit: "string", Category: "Aircraft", Description: "Aircraft model string", ReadOnly: true},
	{Name: "ATC ID", Unit: "string", Category: "Aircraft", Description: "Aircraft tail number / registration", ReadOnly: true},
	{Name: "ATC AIRLINE", Unit: "string", Category: "Aircraft", Description: "Airline name", ReadOnly: true},
	{Name: "ATC FLIGHT NUMBER", Unit: "string", Category: "Aircraft", Description: "Flight number", ReadOnly: true},
	{Name: "ATC HEAVY", Unit: "bool", Category: "Aircraft", Description: "Aircraft is heavy category", ReadOnly: true},
	{Name: "CATEGORY", Unit: "string", Category: "Aircraft", Description: "Aircraft category (e.g. Airplane, Helicopter)", ReadOnly: true},
	{Name: "DESIGN SPEED VS0", Unit: "knots", Category: "Aircraft", Description: "Stall speed in landing config", ReadOnly: true},
	{Name: "DESIGN SPEED VS1", Unit: "knots", Category: "Aircraft", Description: "Stall speed clean", ReadOnly: true},
	{Name: "DESIGN SPEED VC", Unit: "knots", Category: "Aircraft", Description: "Design cruise speed", ReadOnly: true},

	// Position
	{Name: "PLANE LATITUDE", Unit: "degrees", Category: "Position", Description: "Aircraft latitude", ReadOnly: true},
	{Name: "PLANE LONGITUDE", Unit: "degrees", Category: "Position", Description: "Aircraft longitude", ReadOnly: true},
	{Name: "PLANE ALTITUDE", Unit: "feet", Category: "Position", Description: "Aircraft altitude MSL", ReadOnly: true},
	{Name: "PLANE ALT ABOVE GROUND", Unit: "feet", Category: "Position", Description: "Aircraft altitude AGL", ReadOnly: true},
	{Name: "GROUND ALTITUDE", Unit: "feet", Category: "Position", Description: "Ground elevation below aircraft", ReadOnly: true},

	// Speed
	{Name: "AIRSPEED INDICATED", Unit: "knots", Category: "Speed", Description: "Indicated airspeed", ReadOnly: true},
	{Name: "AIRSPEED TRUE", Unit: "knots", Category: "Speed", Description: "True airspeed", ReadOnly: true},
	{Name: "GROUND VELOCITY", Unit: "knots", Category: "Speed", Description: "Ground speed", ReadOnly: true},
	{Name: "VERTICAL SPEED", Unit: "feet per minute", Category: "Speed", Description: "Vertical speed", ReadOnly: true},
	{Name: "AIRSPEED MACH", Unit: "mach", Category: "Speed", Description: "Mach number", ReadOnly: true},

	// Orientation
	{Name: "PLANE HEADING DEGREES TRUE", Unit: "degrees", Category: "Orientation", Description: "True heading", ReadOnly: true},
	{Name: "PLANE HEADING DEGREES MAGNETIC", Unit: "degrees", Category: "Orientation", Description: "Magnetic heading", ReadOnly: true},
	{Name: "HEADING INDICATOR", Unit: "degrees", Category: "Orientation", Description: "Heading indicator reading", ReadOnly: true},
	{Name: "PLANE PITCH DEGREES", Unit: "degrees", Category: "Orientation", Description: "Pitch angle", ReadOnly: true},
	{Name: "PLANE BANK DEGREES", Unit: "degrees", Category: "Orientation", Description: "Bank angle", ReadOnly: true},
	{Name: "INCIDENCE ALPHA", Unit: "degrees", Category: "Orientation", Description: "Angle of attack", ReadOnly: true},

	// Engine
	{Name: "GENERAL ENG RPM:1", Unit: "RPM", Category: "Engine", Description: "Engine 1 RPM", ReadOnly: true},
	{Name: "GENERAL ENG RPM:2", Unit: "RPM", Category: "Engine", Description: "Engine 2 RPM", ReadOnly: true},
	{Name: "ENG N1 RPM:1", Unit: "percent", Category: "Engine", Description: "Engine 1 N1 %", ReadOnly: true},
	{Name: "ENG N1 RPM:2", Unit: "percent", Category: "Engine", Description: "Engine 2 N1 %", ReadOnly: true},
	{Name: "ENG N2 RPM:1", Unit: "percent", Category: "Engine", Description: "Engine 1 N2 %", ReadOnly: true},
	{Name: "ENG N2 RPM:2", Unit: "percent", Category: "Engine", Description: "Engine 2 N2 %", ReadOnly: true},
	{Name: "GENERAL ENG THROTTLE LEVER POSITION:1", Unit: "percent", Category: "Engine", Description: "Throttle lever 1 position", ReadOnly: true},
	{Name: "GENERAL ENG THROTTLE LEVER POSITION:2", Unit: "percent", Category: "Engine", Description: "Throttle lever 2 position", ReadOnly: true},
	{Name: "ENG EXHAUST GAS TEMPERATURE:1", Unit: "rankine", Category: "Engine", Description: "Engine 1 EGT", ReadOnly: true},
	{Name: "ENG OIL PRESSURE:1", Unit: "psf", Category: "Engine", Description: "Engine 1 oil pressure", ReadOnly: true},
	{Name: "ENG OIL TEMPERATURE:1", Unit: "rankine", Category: "Engine", Description: "Engine 1 oil temp", ReadOnly: true},
	{Name: "TURB ENG ITT:1", Unit: "rankine", Category: "Engine", Description: "Engine 1 ITT (turbine)", ReadOnly: true},

	// Fuel
	{Name: "FUEL TOTAL QUANTITY", Unit: "gallons", Category: "Fuel", Description: "Total fuel quantity", ReadOnly: true},
	{Name: "FUEL TOTAL QUANTITY WEIGHT", Unit: "pounds", Category: "Fuel", Description: "Total fuel weight", ReadOnly: true},
	{Name: "FUEL LEFT QUANTITY", Unit: "gallons", Category: "Fuel", Description: "Left tank fuel", ReadOnly: true},
	{Name: "FUEL RIGHT QUANTITY", Unit: "gallons", Category: "Fuel", Description: "Right tank fuel", ReadOnly: true},
	{Name: "ENG FUEL FLOW GPH:1", Unit: "gallons per hour", Category: "Fuel", Description: "Engine 1 fuel flow", ReadOnly: true},
	{Name: "ENG FUEL FLOW GPH:2", Unit: "gallons per hour", Category: "Fuel", Description: "Engine 2 fuel flow", ReadOnly: true},
	{Name: "ESTIMATED FUEL FLOW", Unit: "pounds per hour", Category: "Fuel", Description: "Estimated fuel flow total", ReadOnly: true},

	// Controls
	{Name: "FLAPS HANDLE INDEX", Unit: "number", Category: "Controls", Description: "Flaps lever position index", ReadOnly: true},
	{Name: "FLAPS HANDLE PERCENT", Unit: "percent over 100", Category: "Controls", Description: "Flaps position percent", ReadOnly: true},
	{Name: "TRAILING EDGE FLAPS LEFT ANGLE", Unit: "degrees", Category: "Controls", Description: "Left flap angle", ReadOnly: true},
	{Name: "GEAR HANDLE POSITION", Unit: "bool", Category: "Controls", Description: "Gear handle position (1=down)", ReadOnly: true},
	{Name: "GEAR LEFT POSITION", Unit: "percent over 100", Category: "Controls", Description: "Left gear extension", ReadOnly: true},
	{Name: "GEAR CENTER POSITION", Unit: "percent over 100", Category: "Controls", Description: "Center gear extension", ReadOnly: true},
	{Name: "GEAR RIGHT POSITION", Unit: "percent over 100", Category: "Controls", Description: "Right gear extension", ReadOnly: true},
	{Name: "SPOILERS HANDLE POSITION", Unit: "percent over 100", Category: "Controls", Description: "Spoilers handle position", ReadOnly: true},
	{Name: "SPOILERS ARMED", Unit: "bool", Category: "Controls", Description: "Spoilers armed", ReadOnly: true},
	{Name: "AILERON POSITION", Unit: "position", Category: "Controls", Description: "Aileron deflection (-1 to 1)", ReadOnly: true},
	{Name: "ELEVATOR POSITION", Unit: "position", Category: "Controls", Description: "Elevator deflection (-1 to 1)", ReadOnly: true},
	{Name: "RUDDER POSITION", Unit: "position", Category: "Controls", Description: "Rudder deflection (-1 to 1)", ReadOnly: true},
	{Name: "ELEVATOR TRIM POSITION", Unit: "degrees", Category: "Controls", Description: "Elevator trim angle", ReadOnly: true},
	{Name: "RUDDER TRIM PCT", Unit: "percent over 100", Category: "Controls", Description: "Rudder trim percent", ReadOnly: true},
	{Name: "AILERON TRIM PCT", Unit: "percent over 100", Category: "Controls", Description: "Aileron trim percent", ReadOnly: true},

	// Autopilot
	{Name: "AUTOPILOT MASTER", Unit: "bool", Category: "Autopilot", Description: "Autopilot master switch on", ReadOnly: true},
	{Name: "AUTOPILOT FLIGHT DIRECTOR ACTIVE", Unit: "bool", Category: "Autopilot", Description: "Flight director active", ReadOnly: true},
	{Name: "AUTOPILOT ALTITUDE LOCK", Unit: "bool", Category: "Autopilot", Description: "Altitude hold active", ReadOnly: true},
	{Name: "AUTOPILOT ALTITUDE LOCK VAR", Unit: "feet", Category: "Autopilot", Description: "Selected altitude", ReadOnly: true},
	{Name: "AUTOPILOT HEADING LOCK", Unit: "bool", Category: "Autopilot", Description: "Heading hold active", ReadOnly: true},
	{Name: "AUTOPILOT HEADING LOCK DIR", Unit: "degrees", Category: "Autopilot", Description: "Selected heading", ReadOnly: true},
	{Name: "AUTOPILOT AIRSPEED HOLD", Unit: "bool", Category: "Autopilot", Description: "Speed hold active", ReadOnly: true},
	{Name: "AUTOPILOT AIRSPEED HOLD VAR", Unit: "knots", Category: "Autopilot", Description: "Selected speed", ReadOnly: true},
	{Name: "AUTOPILOT VERTICAL HOLD", Unit: "bool", Category: "Autopilot", Description: "VS hold active", ReadOnly: true},
	{Name: "AUTOPILOT VERTICAL HOLD VAR", Unit: "feet per minute", Category: "Autopilot", Description: "Selected vertical speed", ReadOnly: true},
	{Name: "AUTOPILOT NAV1 LOCK", Unit: "bool", Category: "Autopilot", Description: "NAV hold active", ReadOnly: true},
	{Name: "AUTOPILOT APPROACH HOLD", Unit: "bool", Category: "Autopilot", Description: "Approach mode active", ReadOnly: true},
	{Name: "AUTOPILOT WING LEVELER", Unit: "bool", Category: "Autopilot", Description: "Wing leveler active", ReadOnly: true},
	{Name: "AUTOPILOT THROTTLE ARM", Unit: "bool", Category: "Autopilot", Description: "Autothrottle armed", ReadOnly: true},
	{Name: "AUTOPILOT MACH HOLD", Unit: "bool", Category: "Autopilot", Description: "Mach hold active", ReadOnly: true},

	// Electrical
	{Name: "ELECTRICAL MASTER BATTERY", Unit: "bool", Category: "Electrical", Description: "Master battery on", ReadOnly: true},
	{Name: "GENERAL ENG MASTER ALTERNATOR:1", Unit: "bool", Category: "Electrical", Description: "Alternator 1 on", ReadOnly: true},
	{Name: "AVIONICS MASTER SWITCH", Unit: "bool", Category: "Electrical", Description: "Avionics master on", ReadOnly: true},
	{Name: "ELECTRICAL TOTAL LOAD AMPS", Unit: "amperes", Category: "Electrical", Description: "Total electrical load", ReadOnly: true},
	{Name: "ELECTRICAL BATTERY VOLTAGE", Unit: "volts", Category: "Electrical", Description: "Battery voltage", ReadOnly: true},

	// Lights
	{Name: "LIGHT LANDING", Unit: "bool", Category: "Lights", Description: "Landing lights on", ReadOnly: true},
	{Name: "LIGHT STROBE", Unit: "bool", Category: "Lights", Description: "Strobe lights on", ReadOnly: true},
	{Name: "LIGHT NAV", Unit: "bool", Category: "Lights", Description: "Nav lights on", ReadOnly: true},
	{Name: "LIGHT BEACON", Unit: "bool", Category: "Lights", Description: "Beacon on", ReadOnly: true},
	{Name: "LIGHT TAXI", Unit: "bool", Category: "Lights", Description: "Taxi lights on", ReadOnly: true},
	{Name: "LIGHT CABIN", Unit: "bool", Category: "Lights", Description: "Cabin lights on", ReadOnly: true},

	// Brakes
	{Name: "BRAKE LEFT POSITION", Unit: "position 32k", Category: "Brakes", Description: "Left brake position", ReadOnly: true},
	{Name: "BRAKE RIGHT POSITION", Unit: "position 32k", Category: "Brakes", Description: "Right brake position", ReadOnly: true},
	{Name: "BRAKE PARKING POSITION", Unit: "bool", Category: "Brakes", Description: "Parking brake set", ReadOnly: true},
	{Name: "AUTO BRAKE SWITCH CB", Unit: "number", Category: "Brakes", Description: "Autobrake setting", ReadOnly: true},

	// Radio/Nav
	{Name: "COM ACTIVE FREQUENCY:1", Unit: "MHz", Category: "Radio", Description: "COM1 active frequency", ReadOnly: true},
	{Name: "COM STANDBY FREQUENCY:1", Unit: "MHz", Category: "Radio", Description: "COM1 standby frequency", ReadOnly: true},
	{Name: "NAV ACTIVE FREQUENCY:1", Unit: "MHz", Category: "Radio", Description: "NAV1 active frequency", ReadOnly: true},
	{Name: "TRANSPONDER CODE:1", Unit: "number", Category: "Radio", Description: "Transponder code (squawk)", ReadOnly: true},
	{Name: "GPS GROUND SPEED", Unit: "knots", Category: "Radio", Description: "GPS ground speed", ReadOnly: true},
	{Name: "GPS WP DISTANCE", Unit: "meters", Category: "Radio", Description: "Distance to next GPS waypoint", ReadOnly: true},
	{Name: "GPS WP ETE", Unit: "seconds", Category: "Radio", Description: "ETE to next GPS waypoint", ReadOnly: true},
	{Name: "GPS FLIGHT PLAN WP INDEX", Unit: "number", Category: "Radio", Description: "Current flight plan waypoint index", ReadOnly: true},
	{Name: "GPS FLIGHT PLAN WP COUNT", Unit: "number", Category: "Radio", Description: "Total flight plan waypoints", ReadOnly: true},

	// Environment
	{Name: "AMBIENT TEMPERATURE", Unit: "celsius", Category: "Environment", Description: "Outside air temperature", ReadOnly: true},
	{Name: "AMBIENT WIND VELOCITY", Unit: "knots", Category: "Environment", Description: "Wind speed", ReadOnly: true},
	{Name: "AMBIENT WIND DIRECTION", Unit: "degrees", Category: "Environment", Description: "Wind direction", ReadOnly: true},
	{Name: "BAROMETER PRESSURE", Unit: "millibars", Category: "Environment", Description: "Barometric pressure", ReadOnly: true},
	{Name: "AMBIENT VISIBILITY", Unit: "meters", Category: "Environment", Description: "Visibility", ReadOnly: true},
	{Name: "AMBIENT PRECIP STATE", Unit: "mask", Category: "Environment", Description: "Precipitation state flags", ReadOnly: true},
	{Name: "SEA LEVEL PRESSURE", Unit: "millibars", Category: "Environment", Description: "Sea level pressure (QNH)", ReadOnly: true},

	// Simulation
	{Name: "SIM ON GROUND", Unit: "bool", Category: "Simulation", Description: "Aircraft on ground", ReadOnly: true},
	{Name: "SIM SPEED", Unit: "number", Category: "Simulation", Description: "Simulation rate multiplier", ReadOnly: true},
	{Name: "ABSOLUTE TIME", Unit: "seconds", Category: "Simulation", Description: "Sim time since midnight Jan 1", ReadOnly: true},
	{Name: "ZULU TIME", Unit: "seconds", Category: "Simulation", Description: "Zulu time of day in seconds", ReadOnly: true},
	{Name: "LOCAL TIME", Unit: "seconds", Category: "Simulation", Description: "Local time of day in seconds", ReadOnly: true},

	// G-Forces
	{Name: "G FORCE", Unit: "gforce", Category: "G-Forces", Description: "Current G-force", ReadOnly: true},
	{Name: "MAX G FORCE", Unit: "gforce", Category: "G-Forces", Description: "Maximum G-force experienced", ReadOnly: true},
	{Name: "MIN G FORCE", Unit: "gforce", Category: "G-Forces", Description: "Minimum G-force experienced", ReadOnly: true},

	// Weight
	{Name: "TOTAL WEIGHT", Unit: "pounds", Category: "Weight", Description: "Total aircraft weight", ReadOnly: true},
	{Name: "EMPTY WEIGHT", Unit: "pounds", Category: "Weight", Description: "Empty weight", ReadOnly: true},
	{Name: "MAX GROSS WEIGHT", Unit: "pounds", Category: "Weight", Description: "Maximum gross weight", ReadOnly: true},

	// Mixture/Prop
	{Name: "GENERAL ENG MIXTURE LEVER POSITION:1", Unit: "percent", Category: "Engine", Description: "Mixture lever 1 position", ReadOnly: true},
	{Name: "GENERAL ENG PROPELLER LEVER POSITION:1", Unit: "percent", Category: "Engine", Description: "Prop lever 1 position", ReadOnly: true},
	{Name: "PROP RPM:1", Unit: "RPM", Category: "Engine", Description: "Propeller 1 RPM", ReadOnly: true},

	// Pitot/Anti-ice
	{Name: "PITOT HEAT", Unit: "bool", Category: "Anti-Ice", Description: "Pitot heat on", ReadOnly: true},
	{Name: "PITOT ICE PCT", Unit: "percent over 100", Category: "Anti-Ice", Description: "Pitot tube ice accumulation", ReadOnly: true},
	{Name: "STRUCTURAL ICE PCT", Unit: "percent over 100", Category: "Anti-Ice", Description: "Structural ice accumulation", ReadOnly: true},
	{Name: "STRUCTURAL DEICE SWITCH", Unit: "bool", Category: "Anti-Ice", Description: "Structural de-ice on", ReadOnly: true},
	{Name: "ENG ANTI ICE:1", Unit: "bool", Category: "Anti-Ice", Description: "Engine 1 anti-ice on", ReadOnly: true},
	{Name: "WINDSHIELD DEICE SWITCH", Unit: "bool", Category: "Anti-Ice", Description: "Windshield deice on", ReadOnly: true},

	// Warnings
	{Name: "STALL WARNING", Unit: "bool", Category: "Warnings", Description: "Stall warning active", ReadOnly: true},
	{Name: "OVERSPEED WARNING", Unit: "bool", Category: "Warnings", Description: "Overspeed warning active", ReadOnly: true},
	{Name: "AUTOPILOT DISENGAGED", Unit: "bool", Category: "Warnings", Description: "Autopilot disconnect warning", ReadOnly: true},

	// Hydraulics
	{Name: "HYDRAULIC PRESSURE:1", Unit: "psf", Category: "Hydraulics", Description: "Hydraulic system 1 pressure", ReadOnly: true},
	{Name: "HYDRAULIC PRESSURE:2", Unit: "psf", Category: "Hydraulics", Description: "Hydraulic system 2 pressure", ReadOnly: true},
	{Name: "HYDRAULIC SWITCH:1", Unit: "bool", Category: "Hydraulics", Description: "Hydraulic pump 1 on", ReadOnly: true},
	{Name: "HYDRAULIC SWITCH:2", Unit: "bool", Category: "Hydraulics", Description: "Hydraulic pump 2 on", ReadOnly: true},

	// Pressurization
	{Name: "PRESSURIZATION CABIN ALTITUDE", Unit: "feet", Category: "Pressurization", Description: "Cabin pressure altitude", ReadOnly: true},
	{Name: "PRESSURIZATION DIFF PRESSURE PSI", Unit: "psi", Category: "Pressurization", Description: "Cabin differential pressure", ReadOnly: true},
	{Name: "PRESSURIZATION CABIN ALTITUDE GOAL", Unit: "feet", Category: "Pressurization", Description: "Target cabin altitude", ReadOnly: true},
	{Name: "PRESSURIZATION DUMP SWITCH", Unit: "bool", Category: "Pressurization", Description: "Pressure dump valve open", ReadOnly: true},

	// GPS (expanded)
	{Name: "GPS WP NEXT ID", Unit: "string", Category: "GPS", Description: "Next waypoint identifier", ReadOnly: true},
	{Name: "GPS WP PREV VALID", Unit: "bool", Category: "GPS", Description: "Previous waypoint is valid", ReadOnly: true},
	{Name: "GPS WP TRUE BEARING", Unit: "degrees", Category: "GPS", Description: "True bearing to next waypoint", ReadOnly: true},
	{Name: "GPS WP CROSS TRK", Unit: "meters", Category: "GPS", Description: "Cross-track error", ReadOnly: true},
	{Name: "GPS COURSE TO STEER", Unit: "degrees", Category: "GPS", Description: "Course to steer to next WP", ReadOnly: true},
	{Name: "GPS IS ACTIVE FLIGHT PLAN", Unit: "bool", Category: "GPS", Description: "Flight plan is active", ReadOnly: true},
	{Name: "GPS IS ACTIVE WAY POINT", Unit: "bool", Category: "GPS", Description: "Active waypoint is valid", ReadOnly: true},
	{Name: "GPS OBS ACTIVE", Unit: "bool", Category: "GPS", Description: "GPS OBS mode active", ReadOnly: true},
	{Name: "GPS OBS VALUE", Unit: "degrees", Category: "GPS", Description: "GPS OBS course setting", ReadOnly: true},
	{Name: "GPS DRIVES NAV1", Unit: "bool", Category: "GPS", Description: "GPS is driving NAV1", ReadOnly: true},

	// Radio (expanded)
	{Name: "COM ACTIVE FREQUENCY:2", Unit: "MHz", Category: "Radio", Description: "COM2 active frequency", ReadOnly: true},
	{Name: "COM STANDBY FREQUENCY:2", Unit: "MHz", Category: "Radio", Description: "COM2 standby frequency", ReadOnly: true},
	{Name: "NAV ACTIVE FREQUENCY:2", Unit: "MHz", Category: "Radio", Description: "NAV2 active frequency", ReadOnly: true},
	{Name: "NAV STANDBY FREQUENCY:1", Unit: "MHz", Category: "Radio", Description: "NAV1 standby frequency", ReadOnly: true},
	{Name: "NAV STANDBY FREQUENCY:2", Unit: "MHz", Category: "Radio", Description: "NAV2 standby frequency", ReadOnly: true},
	{Name: "ADF ACTIVE FREQUENCY:1", Unit: "Hz", Category: "Radio", Description: "ADF1 active frequency", ReadOnly: true},
	{Name: "ADF STANDBY FREQUENCY:1", Unit: "Hz", Category: "Radio", Description: "ADF1 standby frequency", ReadOnly: true},
	{Name: "NAV OBS:1", Unit: "degrees", Category: "Radio", Description: "NAV1 OBS setting", ReadOnly: true},
	{Name: "NAV OBS:2", Unit: "degrees", Category: "Radio", Description: "NAV2 OBS setting", ReadOnly: true},
	{Name: "NAV CDI:1", Unit: "number", Category: "Radio", Description: "NAV1 CDI needle deflection", ReadOnly: true},
	{Name: "NAV GSI:1", Unit: "number", Category: "Radio", Description: "NAV1 glideslope needle deflection", ReadOnly: true},
	{Name: "NAV HAS NAV:1", Unit: "bool", Category: "Radio", Description: "NAV1 signal received", ReadOnly: true},
	{Name: "NAV HAS GLIDE SLOPE:1", Unit: "bool", Category: "Radio", Description: "NAV1 glideslope available", ReadOnly: true},
	{Name: "DME DISTANCE:1", Unit: "nautical miles", Category: "Radio", Description: "DME1 distance", ReadOnly: true},
	{Name: "DME SPEED:1", Unit: "knots", Category: "Radio", Description: "DME1 ground speed", ReadOnly: true},

	// Engine (expanded for 4-engine)
	{Name: "GENERAL ENG RPM:3", Unit: "RPM", Category: "Engine", Description: "Engine 3 RPM", ReadOnly: true},
	{Name: "GENERAL ENG RPM:4", Unit: "RPM", Category: "Engine", Description: "Engine 4 RPM", ReadOnly: true},
	{Name: "ENG N1 RPM:3", Unit: "percent", Category: "Engine", Description: "Engine 3 N1 %", ReadOnly: true},
	{Name: "ENG N1 RPM:4", Unit: "percent", Category: "Engine", Description: "Engine 4 N1 %", ReadOnly: true},
	{Name: "GENERAL ENG THROTTLE LEVER POSITION:3", Unit: "percent", Category: "Engine", Description: "Throttle lever 3 position", ReadOnly: true},
	{Name: "GENERAL ENG THROTTLE LEVER POSITION:4", Unit: "percent", Category: "Engine", Description: "Throttle lever 4 position", ReadOnly: true},
	{Name: "ENG COMBUSTION:1", Unit: "bool", Category: "Engine", Description: "Engine 1 combustion active", ReadOnly: true},
	{Name: "ENG COMBUSTION:2", Unit: "bool", Category: "Engine", Description: "Engine 2 combustion active", ReadOnly: true},
	{Name: "ENG FAILED:1", Unit: "bool", Category: "Engine", Description: "Engine 1 has failed", ReadOnly: true},
	{Name: "ENG FAILED:2", Unit: "bool", Category: "Engine", Description: "Engine 2 has failed", ReadOnly: true},
	{Name: "ENG ON FIRE:1", Unit: "bool", Category: "Engine", Description: "Engine 1 on fire", ReadOnly: true},
	{Name: "ENG ON FIRE:2", Unit: "bool", Category: "Engine", Description: "Engine 2 on fire", ReadOnly: true},
	{Name: "GENERAL ENG STARTER:1", Unit: "bool", Category: "Engine", Description: "Engine 1 starter active", ReadOnly: true},
	{Name: "GENERAL ENG STARTER:2", Unit: "bool", Category: "Engine", Description: "Engine 2 starter active", ReadOnly: true},
	{Name: "NUMBER OF ENGINES", Unit: "number", Category: "Engine", Description: "Number of engines on aircraft", ReadOnly: true},

	// Doors/Exits
	{Name: "EXIT OPEN:0", Unit: "percent over 100", Category: "Doors", Description: "Main door open percent", ReadOnly: true},
	{Name: "EXIT OPEN:1", Unit: "percent over 100", Category: "Doors", Description: "Exit 1 open percent", ReadOnly: true},
	{Name: "CANOPY OPEN", Unit: "percent over 100", Category: "Doors", Description: "Canopy open percent", ReadOnly: true},

	// Cabin
	{Name: "CABIN SEATBELTS ALERT SWITCH", Unit: "bool", Category: "Cabin", Description: "Seatbelt sign on", ReadOnly: true},
	{Name: "CABIN NO SMOKING ALERT SWITCH", Unit: "bool", Category: "Cabin", Description: "No smoking sign on", ReadOnly: true},

	// Helicopter
	{Name: "ROTOR BRAKE ACTIVE", Unit: "bool", Category: "Helicopter", Description: "Rotor brake engaged", ReadOnly: true},
	{Name: "ROTOR BRAKE HANDLE POS", Unit: "percent over 100", Category: "Helicopter", Description: "Rotor brake handle position", ReadOnly: true},
	{Name: "ROTOR RPM:1", Unit: "RPM", Category: "Helicopter", Description: "Main rotor RPM", ReadOnly: true},
	{Name: "ROTOR RPM PCT:1", Unit: "percent", Category: "Helicopter", Description: "Main rotor RPM percent of normal", ReadOnly: true},
	{Name: "COLLECTIVE POSITION", Unit: "percent over 100", Category: "Helicopter", Description: "Collective position", ReadOnly: true},
	{Name: "ROTOR GOVERNOR ACTIVE", Unit: "bool", Category: "Helicopter", Description: "Rotor governor active", ReadOnly: true},
	{Name: "ROTOR CLUTCH ACTIVE", Unit: "bool", Category: "Helicopter", Description: "Rotor clutch engaged", ReadOnly: true},

	// Fuel (expanded)
	{Name: "FUEL CENTER QUANTITY", Unit: "gallons", Category: "Fuel", Description: "Center tank fuel", ReadOnly: true},
	{Name: "FUEL SELECTED QUANTITY", Unit: "gallons", Category: "Fuel", Description: "Selected tank fuel quantity", ReadOnly: true},
	{Name: "FUEL CROSS FEED", Unit: "enum", Category: "Fuel", Description: "Crossfeed valve state", ReadOnly: true},
	{Name: "FUEL TANK SELECTOR:1", Unit: "enum", Category: "Fuel", Description: "Fuel tank selector position", ReadOnly: true},

	// Altitude (expanded)
	{Name: "INDICATED ALTITUDE", Unit: "feet", Category: "Position", Description: "Indicated altitude (adjusted for altimeter setting)", ReadOnly: true},
	{Name: "RADIO HEIGHT", Unit: "feet", Category: "Position", Description: "Radar altimeter height AGL", ReadOnly: true},
	{Name: "PRESSURE ALTITUDE", Unit: "feet", Category: "Position", Description: "Pressure altitude (standard setting)", ReadOnly: true},
	{Name: "KOHLSMAN SETTING MB", Unit: "millibars", Category: "Instruments", Description: "Altimeter barometric setting", ReadOnly: true},
	{Name: "DECISION HEIGHT", Unit: "feet", Category: "Instruments", Description: "Decision height setting", ReadOnly: true},
	{Name: "DECISION ALTITUDE MSL", Unit: "feet", Category: "Instruments", Description: "Decision altitude MSL", ReadOnly: true},

	// Surface
	{Name: "SURFACE TYPE", Unit: "enum", Category: "Surface", Description: "Surface type (0=concrete, 1=grass, 2=water, etc.)", ReadOnly: true},
	{Name: "SURFACE CONDITION", Unit: "enum", Category: "Surface", Description: "Surface condition (0=normal, 1=wet, 2=icy, etc.)", ReadOnly: true},
	{Name: "ON ANY RUNWAY", Unit: "bool", Category: "Surface", Description: "Aircraft is on a runway", ReadOnly: true},

	// Marker Beacons
	{Name: "INNER MARKER", Unit: "bool", Category: "Radio", Description: "Inner marker beacon active", ReadOnly: true},
	{Name: "MIDDLE MARKER", Unit: "bool", Category: "Radio", Description: "Middle marker beacon active", ReadOnly: true},
	{Name: "OUTER MARKER", Unit: "bool", Category: "Radio", Description: "Outer marker beacon active", ReadOnly: true},

	// ILS/NAV extended
	{Name: "NAV GS FLAG:1", Unit: "bool", Category: "Radio", Description: "NAV1 glideslope flag (unreliable signal)", ReadOnly: true},
	{Name: "NAV TOFROM:1", Unit: "enum", Category: "Radio", Description: "NAV1 TO/FROM indicator (0=off, 1=to, 2=from)", ReadOnly: true},
	{Name: "NAV TOFROM:2", Unit: "enum", Category: "Radio", Description: "NAV2 TO/FROM indicator", ReadOnly: true},
	{Name: "HSI CDI NEEDLE", Unit: "number", Category: "Radio", Description: "HSI course deviation needle", ReadOnly: true},
	{Name: "HSI GSI NEEDLE", Unit: "number", Category: "Radio", Description: "HSI glideslope needle", ReadOnly: true},

	// Autopilot (expanded)
	{Name: "AUTOPILOT FLIGHT LEVEL CHANGE", Unit: "bool", Category: "Autopilot", Description: "FLC mode active", ReadOnly: true},
	{Name: "AUTOPILOT MANAGED SPEED IN MACH", Unit: "bool", Category: "Autopilot", Description: "Speed hold is in mach mode", ReadOnly: true},
	{Name: "AUTOPILOT MACH HOLD VAR", Unit: "mach", Category: "Autopilot", Description: "Selected mach number", ReadOnly: true},
	{Name: "AUTOPILOT YAW DAMPER", Unit: "bool", Category: "Autopilot", Description: "Yaw damper active", ReadOnly: true},
	{Name: "AUTOPILOT BACKCOURSE HOLD", Unit: "bool", Category: "Autopilot", Description: "Back course hold active", ReadOnly: true},

	// Temperature
	{Name: "TOTAL AIR TEMPERATURE", Unit: "celsius", Category: "Environment", Description: "Total air temperature (TAT)", ReadOnly: true},
	{Name: "STANDARD ATM TEMPERATURE", Unit: "celsius", Category: "Environment", Description: "ISA standard temperature at altitude", ReadOnly: true},

	// Payload/CG
	{Name: "PAYLOAD STATION COUNT", Unit: "number", Category: "Weight", Description: "Number of payload stations", ReadOnly: true},
	{Name: "PAYLOAD STATION WEIGHT:1", Unit: "pounds", Category: "Weight", Description: "Payload station 1 weight", ReadOnly: true},
	{Name: "PAYLOAD STATION WEIGHT:2", Unit: "pounds", Category: "Weight", Description: "Payload station 2 weight", ReadOnly: true},
	{Name: "PAYLOAD STATION WEIGHT:3", Unit: "pounds", Category: "Weight", Description: "Payload station 3 weight", ReadOnly: true},
	{Name: "PAYLOAD STATION WEIGHT:4", Unit: "pounds", Category: "Weight", Description: "Payload station 4 weight", ReadOnly: true},
	{Name: "CG PERCENT", Unit: "percent over 100", Category: "Weight", Description: "Center of gravity position (% MAC)", ReadOnly: true},
	{Name: "CG AFT LIMIT", Unit: "percent over 100", Category: "Weight", Description: "Aft CG limit", ReadOnly: true},
	{Name: "CG FWD LIMIT", Unit: "percent over 100", Category: "Weight", Description: "Forward CG limit", ReadOnly: true},

	// V-Speeds (expanded)
	{Name: "DESIGN SPEED CLIMB", Unit: "knots", Category: "Aircraft", Description: "Design climb speed (Vy)", ReadOnly: true},
	{Name: "DESIGN SPEED MIN ROTATION", Unit: "knots", Category: "Aircraft", Description: "Minimum rotation speed (Vr)", ReadOnly: true},
	{Name: "DESIGN TAKEOFF SPEED", Unit: "knots", Category: "Aircraft", Description: "Design takeoff speed", ReadOnly: true},

	// APU
	{Name: "APU SWITCH", Unit: "bool", Category: "Electrical", Description: "APU master switch on", ReadOnly: true},
	{Name: "APU GENERATOR SWITCH:1", Unit: "bool", Category: "Electrical", Description: "APU generator switch on", ReadOnly: true},
	{Name: "APU PCT RPM", Unit: "percent", Category: "Electrical", Description: "APU RPM percent", ReadOnly: true},
	{Name: "APU VOLTS", Unit: "volts", Category: "Electrical", Description: "APU voltage output", ReadOnly: true},

	// Engine ITT/FF per-engine expanded
	{Name: "TURB ENG ITT:2", Unit: "rankine", Category: "Engine", Description: "Engine 2 ITT (turbine)", ReadOnly: true},
	{Name: "ENG FUEL FLOW PPH:1", Unit: "pounds per hour", Category: "Engine", Description: "Engine 1 fuel flow (lbs/hr)", ReadOnly: true},
	{Name: "ENG FUEL FLOW PPH:2", Unit: "pounds per hour", Category: "Engine", Description: "Engine 2 fuel flow (lbs/hr)", ReadOnly: true},
	{Name: "ENG EXHAUST GAS TEMPERATURE:2", Unit: "rankine", Category: "Engine", Description: "Engine 2 EGT", ReadOnly: true},
	{Name: "ENG OIL PRESSURE:2", Unit: "psf", Category: "Engine", Description: "Engine 2 oil pressure", ReadOnly: true},
	{Name: "ENG OIL TEMPERATURE:2", Unit: "rankine", Category: "Engine", Description: "Engine 2 oil temp", ReadOnly: true},

	// Flight controls axis
	{Name: "AILERON AVERAGE DEFLECTION", Unit: "degrees", Category: "Controls", Description: "Average aileron deflection", ReadOnly: true},
	{Name: "ELEVATOR DEFLECTION", Unit: "degrees", Category: "Controls", Description: "Elevator deflection angle", ReadOnly: true},
	{Name: "RUDDER DEFLECTION", Unit: "degrees", Category: "Controls", Description: "Rudder deflection angle", ReadOnly: true},
}
