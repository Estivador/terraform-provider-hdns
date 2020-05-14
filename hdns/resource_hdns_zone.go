package hdns

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{}
}
