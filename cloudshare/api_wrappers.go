package cloudshare

func (envs *Environments) envByName(name string) *Environment {
	for _, env := range *envs {
		if env.Name == name {
			return &env
		}
	}
	return nil
}

/* GetEnvironmentByName is a convenience function that searches for an environment by name
and return nil if not found */
func (c *Client) GetEnvironmentByName(name string) (*Environment, error) {
	allEnvs := Environments{}
	apierr := c.GetEnvironments(true, "allvisible", &allEnvs)
	if apierr != nil {
		return nil, apierr
	}
	return allEnvs.envByName(name), nil
}

func EnvIDToURL(envID string) string {
	return "https://use.cloudshare.com/Ent/Environment.mvc/View/" + envID[2:]
}

func (e *Environment) URL() string {
	return EnvIDToURL(e.ID)
}
