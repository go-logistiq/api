package sql

const AllClients = `
	SELECT id, slug, name, group_id
	FROM clients`

const GetClientBySlug = `
	SELECT id, slug, name, group_id
	FROM clients
	WHERE group_id = $1 AND slug = $2`
