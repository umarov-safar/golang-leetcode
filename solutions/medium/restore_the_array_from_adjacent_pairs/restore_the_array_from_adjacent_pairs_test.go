package restorefinishinhorder

import (
	"slices"
	"testing"
)

func TestRecoverOrder(t *testing.T) {
    order := []int{3,1,2,5,4}
    friends := []int{1,3,4}
    
    except := []int{3, 1, 4}

    result := RecoverOrder(order, friends)
    if !slices.Equal(result, except) {
        t.Errorf("Not restore correctly")
    }
}