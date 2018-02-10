package models

// GetAdmin is for Login Method.
func GetAdmin() (string, string, error) {
	var (
		name   string
		passwd string
	)
	rows, err := db.Query("SELECT name, passwd FROM users WHERE name=\"Amyang\" limit 1")
	defer rows.Close()
	if err != nil {
		return "", "", err
	}
	for rows.Next() {
		rows.Scan(&name, &passwd)
	}
	return name, passwd, nil
}
