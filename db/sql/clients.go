package sql

const AllClients = `
	SELECT id, slug, name, group_id
	FROM clients`

const GetClientBySlug = `
	SELECT id, slug, name, group_id
	FROM clients
	WHERE slug = $1`
