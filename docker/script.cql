CREATE KEYSPACE IF NOT EXISTS numbersdata WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'cascenter' : 3 };

USE numbersdata;
CREATE TABLE numbers (
    country int, 
    ddd int, 
    number text,
    valid boolean,
    lastView int,
    updatedAt timestamp,
    PRIMARY KEY((country, ddd, number), valid)
);

CREATE INDEX ON numbersdata.numbers(valid);
CREATE INDEX ON numbersdata.numbers(ddd);
