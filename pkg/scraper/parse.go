package scraper

import "github.com/samber/lo"

func (r Response) Normalize() Response {
	response := Response{}

	response.Usages = lo.Map(r.Usages, filterUsage)
	return response
}

func filterUsage(usage Usage, _ int) Usage {
	usage.Definitions = lo.Filter(usage.Definitions, filterDefinition)
	return usage
}

func filterDefinition(definition Definition, _ int) bool {
	return len(definition.Examples) > 0
}
