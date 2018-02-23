package cage

type EPG struct {
  BaseAttributes
}
/* NewTenant creates a new Tenant with the appropriate default values */
func NewEPG(name string, alias string, descr string, belongsTo BaseAttributes) *EPG {
  resourceName := fmt.Sprintf("epg-%s", name)
  attrs := BaseAttributes{
    Name: name,
    NameAlias: alias,
    Description: descr,
    Status: "created",
    ObjectClass: "fvAEPg",
    Status: "created",
    DN: fmt.Sprintf("%s/%s", belongsTo.DN, resourceName),
    RN: resourceName
  }

  e := EPG{BaseAttributes: attrs}
  //Do any additional construction logic specific to the EPG here
  return &e
}
