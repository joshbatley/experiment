# Messager

This creates and publishes events into the message queue(kafka?), all events look and are correctly linked to each other. This is basically a dummy event composer, it has a TPS setting to force high loads and simulate an actual service.

TOOD:
- [X] Compose a event
- [ ] Using random data
  - [X] Cache layer to allow actions with same payment id
  - [ ] With most being successful and a chance of failure
  - [X] All payments have a end event
  - [X] For events type only partial data is required as a consumer pieces them together
- [X] TPS settings
- [ ] Event are sent to some queue