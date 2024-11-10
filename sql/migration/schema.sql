CREATE TABLE
    medication_orders (
        file_number INTEGER NOT NULL,
        nurse_name TEXT,
        ward TEXT,
        bed TEXT,
        medication TEXT,
        uom TEXT,
        request_time DATE NOT NULL,
        nurse_remarks TEXT,
        status TEXT NOT NULL,
        pharmacy_remarks TEXT
    );