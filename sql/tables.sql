use goexpert;

-- Create table products
CREATE TABLE cotation (
	id varchar(255), 
	code varchar(80), 
	code_in varchar(80), 
	name varchar(255), 
	high decimal(10,2), 
	low decimal(10,2), 
	var_bid decimal(10,2), 
	pct_change decimal(10,2), 
	bid decimal(10,2), 
	ask decimal(10,2), 
	created_at timestamp, 
	primary key (id)
);

-- Drop table

DROP TABLE cotation;