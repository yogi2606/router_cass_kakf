	CREATE KEYSPACE  "crm"
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
  
  
  CREATE TABLE IF NOT EXISTS crm.customer (
  customerid TEXT,
  firstname TEXT,
  lastname TEXT,
  PRIMARY KEY(customerid, firstname)
);

CREATE TYPE IF NOT EXISTS crm.contact_details (
  mobilenumber TEXT,
  emailid TEXT,
);
CREATE TYPE IF NOT EXISTS crm.address (
  street TEXT,
  district TEXT,
  state TEXT,
  country TEXT,
);

ALTER TABLE customer
ADD  personaladdr address;

ALTER TABLE customer
ADD officialaddr address;


ALTER TABLE customer
ADD  contactdetails contact_details;