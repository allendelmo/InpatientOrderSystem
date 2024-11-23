CREATE TABLE
    medication_orders (
        order_number INTEGER PRIMARY KEY NOT NULL,
        file_number INTEGER NOT NULL,
        nurse_name TEXT,
        ward TEXT,
        bed TEXT,
        medication TEXT,
        quantity INTEGER,
        uom TEXT,
        request_time DATE NOT NULL,
        nurse_remarks TEXT,
        status TEXT NOT NULL,
        pharmacy_remarks TEXT
    );