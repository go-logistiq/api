package sql

const AllGroups = `
	SELECT id, slug, name
	FROM groups`

const GetGroupBySlug = `
	SELECT id, slug, name
	FROM groups
	WHERE slug = $1`
