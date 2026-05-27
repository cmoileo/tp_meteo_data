| Donnée               | JSON (corrigé)      | XML (corrigé)                                                    |
| -------------------- | ------------------- | ---------------------------------------------------------------- |
| Pays                 | country             | station.@country                                                 |
| Coordonnées          | coordinates         | station.coordinates.@lat / @lon                                  |
| Altitude             | altitude            | station.coordinates.@altitude                                    |
| Modèle de capteur    | device              | station.hardware.@model                                          |
| Température          | temperature_celsius | station.observations[n].observation.measure[@type="temperature"] |
| Condition ciel       | conditions          | station.observations[n].@sky                                     |
| Vent                 | wind                | station.observations[n].observation.wind                         |
| Notes (optionnelles) | notes               | station.observations[n].observation.note                         |