# Experiment

Experiment is repo for me to learn Golang(and other maybe eventually), it is a mono-repo of services 
that mimics a payment provider. 

There is a fake message app that pumps data through queues for "real" services then 
use this data as if it there were real events.

## V1.0 goals
All internal services setup, this will be a MVP where all the data is getting generated and settings and authentication is set correct and all services connect and work together. 


This includes the following repos:
- Auth
- Logger
- Messager
- Settings
- Vault

### Goals of stuff to do:
 - GRPs
 - GraphQL
 - Kafka/Event Driven/messaging
 - DBs Postgres, redis, DynamoDB