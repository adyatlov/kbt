package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/adyatlov/kbt/res"
	"gopkg.in/yaml.v2"
)

func PodsToCsv(r io.Reader, w io.Writer) error {
	decoder := yaml.NewDecoder(r)
	pods := res.Pods{}
	err := decoder.Decode(&pods)
	if err != nil {
		return fmt.Errorf("cannot decode YAML: %w", err)
	}
	csvW := csv.NewWriter(w)
	err = csvW.Write([]string{
		"Namespace",
		"Pod Name",
		"Name",
		"Image",
		"Owner Kind",
		"Owner Name",
		"Memory Limit",
		"CPU Limit",
		"Termination Reason",
		"Restart Count",
	})
	for _, p := range pods.Items {
		statuses := make(map[string]res.ContainerStatus)
		for _, s := range p.Status.ContainerStatuses {
			statuses[s.Name] = s
		}
		for _, c := range p.Spec.Containers {
			s := statuses[c.Name]
			err = csvW.Write([]string{
				p.Metadata.Namespace,
				p.Metadata.Name,
				c.Name,
				c.Image,
				p.Metadata.OwnerReferences[0].Kind,
				p.Metadata.OwnerReferences[0].Name,
				c.Resources.Limits.Memory,
				c.Resources.Limits.CPU,
				s.LastState.Terminated.Reason,
				strconv.Itoa(s.RestartCount),
			})
			if err != nil {
				return fmt.Errorf("cannot write to CSV: %w", err)
			}
		}
	}
	csvW.Flush()
	return nil
}
