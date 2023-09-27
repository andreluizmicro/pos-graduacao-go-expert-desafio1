use goexpert;

-- Create table products
CREATE TABLE cotation (
	id varchar(255), 
	code varchar(80), 
	code_in varchar(80), 
	name varchar(255), 
	high decimal(8,4), 
	low decimal(8,4), 
	var_bid decimal(8,4), 
	pct_change decimal(8,4), 
	bid decimal(8,4), 
	ask decimal(8,4), 
	created_at timestamp, 
	primary key (id)
);

-- Drop table

DROP TABLE cotation;