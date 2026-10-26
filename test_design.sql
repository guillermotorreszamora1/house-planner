BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Supply" (
	"name"	TEXT NOT NULL,
	"type_id"	INTEGER NOT NULL,
	"subtype_id"	TEXT,
	"unit_type_id"	INTEGER,
	"quantity_numerator"	INTEGER NOT NULL,
	"quantity_denominator"	INTEGER NOT NULL,
	"supply_id"	INTEGER NOT NULL,
	PRIMARY KEY("supply_id"),
	FOREIGN KEY("subtype_id") REFERENCES "",
	FOREIGN KEY("type_id") REFERENCES "type"("type_id"),
	FOREIGN KEY("unit_type_id") REFERENCES ""
);
CREATE TABLE IF NOT EXISTS "nutritional_category" (
	"category_id"	INTEGER,
	"name"	INTEGER
);
CREATE TABLE IF NOT EXISTS "subtype" (
	"subtype_id"	INTEGER NOT NULL, --Add primary key
	"subtype"	TEXT
);
CREATE TABLE IF NOT EXISTS "supply_nutritional_category" (
	"supply"	INTEGER,
	"category_id"	INTEGER, -- Add autoincrement
	"weight"	INTEGER,
	PRIMARY KEY("category_id","supply"),
	FOREIGN KEY("supply") REFERENCES ""
);
CREATE TABLE IF NOT EXISTS "type" (
	"type_id"	INTEGER NOT NULL, -- Add autoincrement
	"type"	TEXT,
	PRIMARY KEY("type_id")
);
CREATE TABLE IF NOT EXISTS "units" ( --Add primary key
	"unit_type_id"	INTEGER, -- Add autoincrement
	"unit_name"	TEXT
);
COMMIT;
