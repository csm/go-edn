package types
import "fmt"

type TaggedValue struct {
	Tag string
	Value Value
}

func (this TaggedValue) Equals(other Value) bool {
	switch o := other.(type) {
	case TaggedValue: return this.Tag == o.Tag && this.Value.Equals(o.Value)
	default: return this.Value.Equals(other)
	}
}

func (this TaggedValue) String() string {
	return fmt.Sprint("#", this.Tag, " ", this.Value.String())
}