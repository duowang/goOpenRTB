// OpenRTB 2.1 spec
package openrtb

// Use value for builtin type and pointer for struct

// Bid Request
type Request struct {
	Id      string       `json:"id"`
	Imp     []Impression `json:"imp"`
	Site    *Site        `json:"site"`
	App     *App         `json:"app"`
	Device  *Device      `json:"device"`
	User    *User        `json:"user"`
	At      int          `json:"at"`
	Tmux    int          `json:"tmux"`
	Wseat   []string     `json:"wseat"`
	Allimps int          `json:"allimps"`
	Cur     []string     `json:"cur"`
	Bcat    []string     `json:"bcat"`
	Badv    []string     `json:"badv"`
	Ext     *Ext         `json:"ext"`
}

// Impression
type Impression struct {
	Id                string   `json:"id"`
	Banner            *Banner  `json:"banner"`
	Video             *Video   `json:"video"`
	Displaymanager    string   `json:"displaymanager"`
	Displaymanagerver string   `json:"displaymanagerver"`
	Instl             int      `json:"instl"`
	Tagid             string   `json:"tagid"`
	Bidfloor          float32  `json:"bidfloor"`
	Bidfloorcur       string   `json:"bidfloorcur"`
	Iframebuster      []string `json:"iframebuster"`
	Ext               *Ext     `json:"ext"`
}

// Banner
type Banner struct {
	W        int      `json:"w"`
	H        int      `json:"h"`
	Id       string   `json:"id"`
	Pos      int      `json:"pos"`
	Btype    []int    `json:"btype"`
	Battr    []int    `json:"battr"`
	Mimes    []string `json:"mimes"`
	Topframe int      `json:"topframe"`
	Expdir   []int    `json:"expdir"`
	Api      []int    `json:"api"`
	Ext      *Ext     `json:"ext"`
}

// Video
type Video struct {
	Mimes          []string `json:"mimes"`
	Linearity      int      `json:"linearity"`
	Minduration    int      `json:"minduration"`
	Maxduration    int      `json:"maxduration"`
	Protocal       int      `json:"protocal"`
	W              int      `json:"w"`
	H              int      `json:"h"`
	Startdelay     int      `json:"startdelay"`
	Sequence       int      `json:"sequence"`
	Battr          []int    `json:"battr"`
	Type           []string `json:"type"`
	Maxextended    int      `json:"maxextended"`
	Minbitrate     int      `json:"minbitrate"`
	Maxbitrate     int      `json:"minbitrate"`
	Boxingallowed  int      `json:"boxingallowed"`
	Playbackmethod []int    `json:"playbackmethod"`
	Delivery       []int    `json:"delivery"`
	Pos            int      `json:"pos"`
	Api            []int    `json:"api"`
	Ext            *Ext     `json:"ext"`
	// Companioned not supported
	// Companiontype not supported
}

// Site
type Site struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Domain        string     `json:"domain"`
	Cat           []string   `json:"cat"`
	Sectioncat    []string   `json:"sectioncat"`
	Pagecat       []string   `json:"pagecat"`
	Page          string     `json:"page"`
	Privacypolicy int        `json:"privacypolicy"`
	Ref           string     `json:"ref"`
	Search        string     `json:"search"`
	Publisher     *Publisher `json:"publisher"`
	Content       *Content   `json:"content"`
	Keywords      string     `json:"keywords"`
	Ext           *Ext       `json:"ext"`
}

// App
type App struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Domain        string     `json:"domain"`
	Cat           []string   `json:"cat"`
	Sectioncat    []string   `json:"sectioncat"`
	Pagecat       []string   `json:"pagecat"`
	Ver           string     `json:"ver"`
	Bundle        string     `json:"bundle"`
	Privacypolicy int        `json:"privacypolicy"`
	Paid          int        `json:"paid"`
	Publisher     *Publisher `json:"publisher"`
	Content       *Content   `json:"content"`
	Keywords      string     `json:"keywords"`
	Storeurl      string     `json:"storeurl"`
	Ext           *Ext       `json:"ext"`
}

// Content
type Content struct {
	Id                 string    `json:"id"`
	Episode            int       `json:"episode"`
	Title              string    `json:"title"`
	Series             string    `json:"series"`
	Season             string    `json:"season"`
	Url                string    `json:"url"`
	Cat                []string  `json:"cat"`
	Videoquality       int       `json:"videoquality"`
	Keywords           string    `json:"keywords"`
	Contentrating      string    `json:"contentrating"`
	Userrating         string    `json:"userrating"`
	Context            string    `json:"context"`
	Livestream         int       `json:"livestream"`
	Sourcerelationship int       `json:"sourcerelationship"`
	Producer           *Producer `json:"producer"`
	Len                int       `json:"len"`
	QAGmediarating     int       `json:"qagmediarating"`
	Embeddable         int       `json:"embeddable"`
	Language           string    `json:"language"`
	Ext                *Ext      `json:"ext"`
}

// Publisher
type Publisher struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Cat    []string `json:"cat"`
	Domain string   `json:"domain"`
	Ext    *Ext     `json:"ext"`
}

// Producer
type Producer struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Cat    []string `json:"cat"`
	Domain string   `json:"domain"`
	Ext    *Ext     `json:"ext"`
}

// Device
type Device struct {
	Dnt            int    `json:"dnt"`
	Ua             string `json:"ua"`
	Ip             string `json:"ip"`
	Geo            *Geo   `json:"geo"`
	Didsha1        string `json:"didsha1"`
	Didmd5         string `json:"didmd5"`
	Dpidsha1       string `json:"dpidsha1"`
	Dpidmd5        string `json:"dpidmd5"`
	Ipv6           string `json:"ipv6"`
	Carrier        string `json:"carrier"`
	Language       string `json:"language"`
	Make           string `json:"make"`
	Model          string `json:"model"`
	Os             string `json:"os"`
	Osv            string `json:"osv"`
	Js             int    `json:"js"`
	Connectiontype int    `json:"connectiontype"`
	Devicetype     int    `json:"devicetype"`
	Flashver       string `json:"flashver"`
	Ext            *Ext   `json:"ext"`
}

// Geo
type Geo struct {
	Lat           float32 `json:"lat"`
	Lon           float32 `json:"lon"`
	Country       string  `json:"country"`
	Region        string  `json:"region"`
	Regionfips104 string  `json:"regionfips104"`
	Metro         string  `json:"metro"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Type          int     `json:"type"`
	Ext           *Ext    `json:"ext"`
}

// User
type User struct {
	Id         string `json:"id"`
	Buyeruid   string `json:"buyeruid"`
	Yob        int    `json:"yob"`
	Gender     string `json:"gender"`
	Keywords   string `json:"keywords"`
	Customdata string `json:"customdata"`
	Geo        *Geo   `json:"geo"`
	Data       []Data `json:"data"`
	Ext        *Ext   `json:"ext"`
}

// Data
type Data struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Segment []Segment `json:"segment"`
	Ext     *Ext      `json:"ext"`
}

// Segment
type Segment struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Ext   *Ext   `json:"ext"`
}

type Ext map[string]interface{}
