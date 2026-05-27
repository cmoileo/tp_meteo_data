# Change Log

## [0.2]

### Added
- XML parser (`LoadFromXML`) with measure type/pollutant name dispatching and optional note handling

## [0.1]

### Added
- Unified internal model (`model.go`) with Station, Observation, Coordinates, AirQuality
- JSON mirror structs (`jsonsource.go`) with json tags replicating the raw JSON schema
- JSON parser (`LoadFromJSON`) with country name-to-ISO mapping (14 entries)
- XML mirror structs (`xmlsource.go`) with xml tags replicating the raw XML schema