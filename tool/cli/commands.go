// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ComingServer": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/TinyKitten/ComingServer/design
// --out=$(GOPATH)/src/github.com/TinyKitten/ComingServer
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TinyKitten/ComingServer/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// AuthAuthCommand is the command line data structure for the auth action of auth
	AuthAuthCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// AddPeersCommand is the command line data structure for the add action of peers
	AddPeersCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// CurrentLocationPeersCommand is the command line data structure for the current location action of peers
	CurrentLocationPeersCommand struct {
		// ピアID
		ID          int
		PrettyPrint bool
	}

	// ListPeersCommand is the command line data structure for the list action of peers
	ListPeersCommand struct {
		// 最大件数
		Limit int
		// 取得開始位置
		Offset      int
		PrettyPrint bool
	}

	// LocationsPeersCommand is the command line data structure for the locations action of peers
	LocationsPeersCommand struct {
		// ピアID
		ID          int
		PrettyPrint bool
	}

	// RegenerateTokenPeersCommand is the command line data structure for the regenerate token action of peers
	RegenerateTokenPeersCommand struct {
		// ピアID
		ID          int
		PrettyPrint bool
	}

	// SendLocationPeersCommand is the command line data structure for the send location action of peers
	SendLocationPeersCommand struct {
		Payload     string
		ContentType string
		// ピアID
		ID          int
		PrettyPrint bool
	}

	// ShowPeersCommand is the command line data structure for the show action of peers
	ShowPeersCommand struct {
		// ピアID
		ID          int
		PrettyPrint bool
	}

	// UpdatePeersCommand is the command line data structure for the update action of peers
	UpdatePeersCommand struct {
		Payload     string
		ContentType string
		ID          string
		// ピアID
		Code        int
		PrettyPrint bool
	}

	// AddPodsCommand is the command line data structure for the add action of pods
	AddPodsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListPodsCommand is the command line data structure for the list action of pods
	ListPodsCommand struct {
		// 最大件数
		Limit int
		// 取得開始位置
		Offset      int
		PrettyPrint bool
	}

	// RegenerateTokenPodsCommand is the command line data structure for the regenerate token action of pods
	RegenerateTokenPodsCommand struct {
		// ポッドID
		ID          int
		PrettyPrint bool
	}

	// ShowPodsCommand is the command line data structure for the show action of pods
	ShowPodsCommand struct {
		// ポッドID
		ID          int
		PrettyPrint bool
	}

	// UpdatePodsCommand is the command line data structure for the update action of pods
	UpdatePodsCommand struct {
		Payload     string
		ContentType string
		// ポッドID
		ID          int
		PrettyPrint bool
	}

	// AddUsersCommand is the command line data structure for the add action of users
	AddUsersCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListUsersCommand is the command line data structure for the list action of users
	ListUsersCommand struct {
		// 最大件数
		Limit int
		// 取得開始位置
		Offset      int
		PrettyPrint bool
	}

	// ShowUsersCommand is the command line data structure for the show action of users
	ShowUsersCommand struct {
		// ユーザーID
		ID          int
		PrettyPrint bool
	}

	// SendCurrentPeerLocationWebsocketCommand is the command line data structure for the send current peer location action of websocket
	SendCurrentPeerLocationWebsocketCommand struct {
		// 緯度
		Latitude string
		// スクリーンネーム
		Longitude string
		// Secret token for send location
		Token       string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add",
		Short: `add action`,
	}
	tmp1 := new(AddPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers"]`,
		Short: ``,
		Long: `

Payload example:

{
   "code": "TS"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(AddPodsCommand)
	sub = &cobra.Command{
		Use:   `pods ["/v1/pods"]`,
		Short: ``,
		Long: `

Payload example:

{
   "code": "SHIBUYA",
   "latitude": 35.658034,
   "longitude": 139.701636
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(AddUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users"]`,
		Short: ``,
		Long: `

Payload example:

{
   "password": "password",
   "privilege_id": 0,
   "screen_name": "tinykitten"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "auth",
		Short: `JWTトークンを生成する`,
	}
	tmp4 := new(AuthAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/v1/auth"]`,
		Short: ``,
		Long: `

Payload example:

{
   "id": 0,
   "password": "password",
   "screen_name": "tinykitten"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "current-location",
		Short: `ピアの現在位置`,
	}
	tmp5 := new(CurrentLocationPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID/location"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp6 := new(ListPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(ListPodsCommand)
	sub = &cobra.Command{
		Use:   `pods ["/v1/pods"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(ListUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "locations",
		Short: `ピアの位置履歴`,
	}
	tmp9 := new(LocationsPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID/locations"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "regenerate-token",
		Short: `regenerateToken action`,
	}
	tmp10 := new(RegenerateTokenPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID/token"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(RegenerateTokenPodsCommand)
	sub = &cobra.Command{
		Use:   `pods ["/v1/pods/ID/token"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "send-current-peer-location",
		Short: `echo websocket server`,
	}
	tmp12 := new(SendCurrentPeerLocationWebsocketCommand)
	sub = &cobra.Command{
		Use:   `websocket ["/v1/echo"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "send-location",
		Short: `ピア位置送信`,
	}
	tmp13 := new(SendLocationPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "latitude": 35.658034,
   "longitude": 139.701636
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp14 := new(ShowPeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp15 := new(ShowPodsCommand)
	sub = &cobra.Command{
		Use:   `pods ["/v1/pods/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp16 := new(ShowUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp17 := new(UpdatePeersCommand)
	sub = &cobra.Command{
		Use:   `peers ["/v1/peers/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "name": "TS-IPHONE"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp17.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp18 := new(UpdatePodsCommand)
	sub = &cobra.Command{
		Use:   `pods ["/v1/pods/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "latitude": 35.689592,
   "longitude": 139.700413,
   "name": "SHINJUKU"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp18.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swaggerui/") {
		fnd = c.DownloadSwaggerui
		rpath = rpath[11:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the AuthAuthCommand command.
func (cmd *AuthAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/auth"
	}
	var payload client.AuthAuthPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AuthAuth(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AuthAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the AddPeersCommand command.
func (cmd *AddPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/peers"
	}
	var payload client.AddPeersPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddPeers(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the CurrentLocationPeersCommand command.
func (cmd *CurrentLocationPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v/location", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CurrentLocationPeers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CurrentLocationPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ピアID`)
}

// Run makes the HTTP request corresponding to the ListPeersCommand command.
func (cmd *ListPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/peers"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListPeers(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 100, `最大件数`)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, `取得開始位置`)
}

// Run makes the HTTP request corresponding to the LocationsPeersCommand command.
func (cmd *LocationsPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v/locations", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.LocationsPeers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LocationsPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ピアID`)
}

// Run makes the HTTP request corresponding to the RegenerateTokenPeersCommand command.
func (cmd *RegenerateTokenPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v/token", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RegenerateTokenPeers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RegenerateTokenPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ピアID`)
}

// Run makes the HTTP request corresponding to the SendLocationPeersCommand command.
func (cmd *SendLocationPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v", cmd.ID)
	}
	var payload client.SendLocationPeersPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SendLocationPeers(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SendLocationPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ピアID`)
}

// Run makes the HTTP request corresponding to the ShowPeersCommand command.
func (cmd *ShowPeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowPeers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowPeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ピアID`)
}

// Run makes the HTTP request corresponding to the UpdatePeersCommand command.
func (cmd *UpdatePeersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/peers/%v", url.QueryEscape(cmd.ID))
	}
	var payload client.UpdatePeersPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdatePeers(ctx, path, &payload, intFlagVal("code", cmd.Code), cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdatePeersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
	var code int
	cc.Flags().IntVar(&cmd.Code, "code", code, `ピアID`)
}

// Run makes the HTTP request corresponding to the AddPodsCommand command.
func (cmd *AddPodsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/pods"
	}
	var payload client.AddPodsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddPods(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddPodsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListPodsCommand command.
func (cmd *ListPodsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/pods"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListPods(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListPodsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 100, `最大件数`)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, `取得開始位置`)
}

// Run makes the HTTP request corresponding to the RegenerateTokenPodsCommand command.
func (cmd *RegenerateTokenPodsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/pods/%v/token", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RegenerateTokenPods(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RegenerateTokenPodsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ポッドID`)
}

// Run makes the HTTP request corresponding to the ShowPodsCommand command.
func (cmd *ShowPodsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/pods/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowPods(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowPodsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ポッドID`)
}

// Run makes the HTTP request corresponding to the UpdatePodsCommand command.
func (cmd *UpdatePodsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/pods/%v", cmd.ID)
	}
	var payload client.UpdatePodsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdatePods(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdatePodsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ポッドID`)
}

// Run makes the HTTP request corresponding to the AddUsersCommand command.
func (cmd *AddUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/users"
	}
	var payload client.AddUsersPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddUsers(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListUsersCommand command.
func (cmd *ListUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/users"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUsers(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 100, `最大件数`)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, `取得開始位置`)
}

// Run makes the HTTP request corresponding to the ShowUsersCommand command.
func (cmd *ShowUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/users/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUsers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `ユーザーID`)
}

// Run establishes a websocket connection for the SendCurrentPeerLocationWebsocketCommand command.
func (cmd *SendCurrentPeerLocationWebsocketCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/echo"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp19 *float64
	if cmd.Latitude != "" {
		var err error
		tmp19, err = float64Val(cmd.Latitude)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *float64 value", "flag", "--latitude", "err", err)
			return err
		}
	}
	if tmp19 == nil {
		goa.LogError(ctx, "required flag is missing", "flag", "--latitude")
		return fmt.Errorf("required flag latitude is missing")
	}
	var tmp20 *float64
	if cmd.Longitude != "" {
		var err error
		tmp20, err = float64Val(cmd.Longitude)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *float64 value", "flag", "--longitude", "err", err)
			return err
		}
	}
	if tmp20 == nil {
		goa.LogError(ctx, "required flag is missing", "flag", "--longitude")
		return fmt.Errorf("required flag longitude is missing")
	}
	ws, err := c.SendCurrentPeerLocationWebsocket(ctx, path, *tmp19, *tmp20, cmd.Token)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SendCurrentPeerLocationWebsocketCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var latitude string
	cc.Flags().StringVar(&cmd.Latitude, "latitude", latitude, `緯度`)
	var longitude string
	cc.Flags().StringVar(&cmd.Longitude, "longitude", longitude, `スクリーンネーム`)
	var token string
	cc.Flags().StringVar(&cmd.Token, "token", token, `Secret token for send location`)
}
