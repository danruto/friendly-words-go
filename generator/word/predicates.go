package word

var WORDLIST_PREDICATES = []string{
	"aback",
	"abalone",
	"abiding",
	"ablaze",
	"able",
	"aboard",
	"abounding",
	"abrasive",
	"abrupt",
	"absorbed",
	"absorbing",
	"abstracted",
	"abundance",
	"abundant",
	"abyssinian",
	"accessible",
	"accidental",
	"accurate",
	"achieved",
	"acidic",
	"acoustic",
	"actually",
	"acute",
	"adaptable",
	"adaptive",
	"adhesive",
	"adjoining",
	"admitted",
	"adorable",
	"adventurous",
	"aeolian",
	"aerial",
	"agate",
	"aged",
	"agreeable",
	"ahead",
	"airy",
	"ajar",
	"alabaster",
	"alder",
	"alert",
	"alike",
	"alive",
	"alkaline",
	"alluring",
	"almond",
	"almondine",
	"aloud",
	"alpine",
	"aluminum",
	"amazing",
	"ambiguous",
	"ambitious",
	"amenable",
	"amethyst",
	"amplified",
	"amused",
	"amusing",
	"ancient",
	"angry",
	"animated",
	"antique",
	"apple",
	"apricot",
	"aquamarine",
	"aquatic",
	"aromatic",
	"arrow",
	"artistic",
	"ash",
	"aspiring",
	"assorted",
	"astonishing",
	"atlantic",
	"atom",
	"attractive",
	"auspicious",
	"automatic",
	"autumn",
	"available",
	"awake",
	"aware",
	"awesome",
	"axiomatic",
	"azure",
	"balanced",
	"bald",
	"ballistic",
	"balsam",
	"band",
	"basalt",
	"battle",
	"bead",
	"beaded",
	"beautiful",
	"bedecked",
	"befitting",
	"bejewled",
	"believed",
	"bemused",
	"beneficial",
	"berry",
	"best",
	"better",
	"bevel",
	"big",
	"billowy",
	"bird",
	"bitter",
	"bittersweet",
	"blend",
	"bloom",
	"blossom",
	"blue",
	"blush",
	"blushing",
	"boatneck",
	"boiled",
	"boiling",
	"bold",
	"bolder",
	"boom",
	"booming",
	"bottled",
	"bottlenose",
	"boulder",
	"bouncy",
	"boundless",
	"bow",
	"brainy",
	"bramble",
	"branch",
	"branched",
	"brash",
	"brass",
	"brassy",
	"brave",
	"brawny",
	"brazen",
	"breezy",
	"brick",
	"brief",
	"bright",
	"brindle",
	"bristle",
	"broad",
	"broadleaf",
	"broken",
	"bronze",
	"bronzed",
	"brook",
	"bubble",
	"bubbly",
	"bumpy",
	"burly",
	"bustling",
	"busy",
	"butter",
	"buttercup",
	"buttered",
	"butternut",
	"buttery",
	"button",
	"buttoned",
	"bygone",
	"cactus",
	"cake",
	"calico",
	"calm",
	"camp",
	"canary",
	"candied",
	"candle",
	"candy",
	"canyon",
	"capable",
	"capricious",
	"caramel",
	"carbonated",
	"carefree",
	"careful",
	"caring",
	"carnation",
	"carnelian",
	"carpal",
	"casual",
	"cat",
	"caterwauling",
	"catkin",
	"catnip",
	"cautious",
	"cedar",
	"celestial",
	"certain",
	"cerulean",
	"chain",
	"chalk",
	"chambray",
	"changeable",
	"charm",
	"charmed",
	"charming",
	"chartreuse",
	"chatter",
	"checker",
	"checkered",
	"cheddar",
	"cheerful",
	"chemical",
	"cherry",
	"chestnut",
	"chief",
	"childish",
	"childlike",
	"chill",
	"chip",
	"chipped",
	"chisel",
	"chiseled",
	"chivalrous",
	"chlorinated",
	"chocolate",
	"chrome",
	"circular",
	"citrine",
	"clammy",
	"clarity",
	"classic",
	"classy",
	"clean",
	"clear",
	"clever",
	"cliff",
	"climbing",
	"closed",
	"cloud",
	"cloudy",
	"clover",
	"clumsy",
	"coal",
	"cobalt",
	"coconut",
	"coffee",
	"coherent",
	"cold",
	"colorful",
	"colossal",
	"comet",
	"comfortable",
	"common",
	"complete",
	"complex",
	"concise",
	"concrete",
	"confirmed",
	"confused",
	"confusion",
	"congruous",
	"conscious",
	"continuous",
	"cooing",
	"cooked",
	"cookie",
	"cool",
	"cooperative",
	"coordinated",
	"copper",
	"copy",
	"coral",
	"cord",
	"corner",
	"cosmic",
	"cotton",
	"cottony",
	"courageous",
	"crawling",
	"creative",
	"crimson",
	"crocus",
	"crystal",
	"crystalline",
	"cubic",
	"cuboid",
	"cuddly",
	"cultivate",
	"cultured",
	"cumbersome",
	"curious",
	"curse",
	"curved",
	"cut",
	"cyan",
	"cyber",
	"cyclic",
	"cypress",
	"daffodil",
	"daffy",
	"daily",
	"dandelion",
	"dandy",
	"dapper",
	"dark",
	"darkened",
	"dashing",
	"dawn",
	"dazed",
	"dazzling",
	"deadpan",
	"dear",
	"debonair",
	"deciduous",
	"decisive",
	"decorous",
	"dedicated",
	"deep",
	"deeply",
	"defiant",
	"delicate",
	"delicious",
	"delightful",
	"delirious",
	"deluxe",
	"denim",
	"dent",
	"dented",
	"descriptive",
	"desert",
	"deserted",
	"destiny",
	"detailed",
	"determined",
	"developing",
	"diagnostic",
	"diamond",
	"different",
	"difficult",
	"diligent",
	"dirt",
	"disco",
	"discovered",
	"discreet",
	"distinct",
	"dog",
	"dolomite",
	"dorian",
	"dot",
	"dour",
	"dramatic",
	"dull",
	"dune",
	"dust",
	"dusty",
	"dynamic",
	"eager",
	"early",
	"earthy",
	"east",
	"eastern",
	"easy",
	"economic",
	"educated",
	"efficacious",
	"efficient",
	"eggplant",
	"eight",
	"elastic",
	"elated",
	"elderly",
	"electric",
	"elegant",
	"elemental",
	"elite",
	"ember",
	"emerald",
	"eminent",
	"emphasized",
	"empty",
	"enchanted",
	"enchanting",
	"encouraging",
	"endurable",
	"energetic",
	"enormous",
	"enshrined",
	"entertaining",
	"enthusiastic",
	"equable",
	"equal",
	"equatorial",
	"equinox",
	"erratic",
	"estimated",
	"ethereal",
	"evanescent",
	"even",
	"evening",
	"evergreen",
	"everlasting",
	"excellent",
	"excessive",
	"excited",
	"exciting",
	"exclusive",
	"expensive",
	"experienced",
	"extreme",
	"exuberant",
	"exultant",
	"fabulous",
	"faceted",
	"factual",
	"faint",
	"fair",
	"faithful",
	"fallacious",
	"false",
	"familiar",
	"famous",
	"fan",
	"fanatical",
	"fancy",
	"fantastic",
	"fantasy",
	"far",
	"fascinated",
	"fast",
	"fate",
	"fearless",
	"feather",
	"feline",
	"fern",
	"festive",
	"few",
	"field",
	"fierce",
	"fifth",
	"fine",
	"fir",
	"fire",
	"first",
	"fish",
	"five",
	"fixed",
	"flame",
	"flannel",
	"flash",
	"flashy",
	"flat",
	"flawless",
	"flax",
	"flaxen",
	"flicker",
	"flint",
	"florentine",
	"flossy",
	"flourish",
	"flower",
	"flowery",
	"fluff",
	"fluffy",
	"fluorescent",
	"fluoridated",
	"fluttering",
	"flying",
	"foam",
	"foamy",
	"fog",
	"foggy",
	"foil",
	"foregoing",
	"foremost",
	"forest",
	"forested",
	"fork",
	"fortunate",
	"fortune",
	"fossil",
	"foul",
	"four",
	"fourth",
	"fragrant",
	"freckle",
	"free",
	"freezing",
	"frequent",
	"fresh",
	"friendly",
	"frill",
	"fringe",
	"frost",
	"frosted",
	"fuchsia",
	"full",
	"functional",
	"funny",
	"furtive",
	"future",
	"futuristic",
	"gainful",
	"galvanized",
	"gamy",
	"garnet",
	"garrulous",
	"gaudy",
	"gelatinous",
	"gem",
	"general",
	"generated",
	"gentle",
	"geode",
	"giant",
	"giddy",
	"gifted",
	"gigantic",
	"gilded",
	"ginger",
	"glacier",
	"glamorous",
	"glass",
	"glaze",
	"gleaming",
	"glib",
	"glimmer",
	"glistening",
	"glitter",
	"glittery",
	"global",
	"glorious",
	"glory",
	"glossy",
	"glow",
	"glowing",
	"gold",
	"golden",
	"goldenrod",
	"good",
	"goofy",
	"gorgeous",
	"gossamer",
	"graceful",
	"grand",
	"grandiose",
	"granite",
	"grape",
	"grass",
	"grateful",
	"gratis",
	"grave",
	"gravel",
	"gray",
	"great",
	"green",
	"gregarious",
	"grey",
	"grizzled",
	"grizzly",
	"groovy",
	"grove",
	"guiltless",
	"gusty",
	"guttural",
	"habitual",
	"hail",
	"half",
	"hallowed",
	"halved",
	"hammerhead",
	"handsome",
	"handsomely",
	"handy",
	"happy",
	"hardly",
	"harmless",
	"harmonious",
	"harsh",
	"harvest",
	"heady",
	"healthy",
	"heartbreaking",
	"heather",
	"heathered",
	"heavenly",
	"heavy",
	"held",
	"heliotrope",
	"helix",
	"helpful",
	"hexagonal",
	"hickory",
	"highfalutin",
	"highly",
	"hilarious",
	"hill",
	"hip",
	"hissing",
	"historical",
	"holistic",
	"hollow",
	"holy",
	"honey",
	"honeysuckle",
	"honorable",
	"honored",
	"horn",
	"horse",
	"hospitable",
	"hot",
	"hulking",
	"humane",
	"humble",
	"humdrum",
	"humorous",
	"hungry",
	"hurricane",
	"hushed",
	"hyper",
	"hypnotic",
	"iced",
	"icy",
	"illustrious",
	"imaginary",
	"immediate",
	"immense",
	"imminent",
	"impartial",
	"important",
	"imported",
	"impossible",
	"incandescent",
	"inconclusive",
	"incongruous",
	"incredible",
	"indecisive",
	"indigo",
	"indispensable",
	"industrious",
	"inexpensive",
	"infrequent",
	"ink",
	"inky",
	"innate",
	"innovative",
	"inquisitive",
	"insidious",
	"instinctive",
	"intelligent",
	"interesting",
	"intermediate",
	"internal",
	"intriguing",
	"invented",
	"invincible",
	"invited",
	"iodized",
	"ionian",
	"ionized",
	"iridescent",
	"iris",
	"iron",
	"irradiated",
	"island",
	"ivy",
	"jagged",
	"jasper",
	"jazzy",
	"jealous",
	"jelly",
	"jet",
	"jewel",
	"jeweled",
	"jolly",
	"joyous",
	"judicious",
	"jumbled",
	"jumpy",
	"jungle",
	"juniper",
	"just",
	"juvenile",
	"kaput",
	"keen",
	"kind",
	"kindhearted",
	"kindly",
	"kiwi",
	"knotty",
	"knowing",
	"knowledgeable",
	"lace",
	"laced",
	"lackadaisical",
	"lacy",
	"lake",
	"languid",
	"lapis",
	"laser",
	"lateral",
	"lava",
	"lavender",
	"lavish",
	"lead",
	"leaf",
	"lean",
	"learned",
	"leather",
	"leeward",
	"legend",
	"legendary",
	"lemon",
	"level",
	"liberating",
	"light",
	"lightning",
	"like",
	"likeable",
	"lilac",
	"lime",
	"linen",
	"literate",
	"lithe",
	"little",
	"lively",
	"living",
	"lizard",
	"local",
	"locrian",
	"lofty",
	"long",
	"longhaired",
	"longing",
	"lopsided",
	"loud",
	"lovely",
	"loving",
	"low",
	"lowly",
	"luck",
	"lucky",
	"ludicrous",
	"lumbar",
	"luminous",
	"lunar",
	"lush",
	"luxuriant",
	"luxurious",
	"lydian",
	"lying",
	"lyrical",
	"maddening",
	"magenta",
	"magic",
	"magical",
	"magnetic",
	"magnificent",
	"mahogany",
	"maize",
	"majestic",
	"malachite",
	"malleable",
	"mammoth",
	"mango",
	"mangrove",
	"maple",
	"marble",
	"marbled",
	"marked",
	"marmalade",
	"maroon",
	"marred",
	"married",
	"marsh",
	"marshy",
	"marvelous",
	"massive",
	"material",
	"materialistic",
	"mature",
	"maze",
	"meadow",
	"medieval",
	"mellow",
	"melodic",
	"melodious",
	"melted",
	"meowing",
	"merciful",
	"mercurial",
	"mercury",
	"mesquite",
	"metal",
	"meteor",
	"mewing",
	"mica",
	"midi",
	"midnight",
	"mighty",
	"mild",
	"mildly",
	"military",
	"mini",
	"miniature",
	"mint",
	"mirage",
	"mire",
	"mirror",
	"misty",
	"mixed",
	"mixolydian",
	"modern",
	"modest",
	"momentous",
	"moored",
	"morning",
	"motley",
	"mountain",
	"mountainous",
	"mousy",
	"mud",
	"muddled",
	"muddy",
	"mulberry",
	"mutual",
	"mysterious",
	"narrow",
	"natural",
	"navy",
	"near",
	"neat",
	"nebula",
	"nebulous",
	"necessary",
	"neighborly",
	"neon",
	"nervous",
	"nettle",
	"nice",
	"nickel",
	"nifty",
	"night",
	"nimble",
	"nine",
	"ninth",
	"noble",
	"noiseless",
	"nonchalant",
	"nonstop",
	"noon",
	"north",
	"northern",
	"nostalgic",
	"nosy",
	"notch",
	"nova",
	"numerous",
	"nutritious",
	"oasis",
	"observant",
	"obsidian",
	"obtainable",
	"obvious",
	"occipital",
	"oceanic",
	"octagonal",
	"odd",
	"oil",
	"olive",
	"olivine",
	"omniscient",
	"onyx",
	"opalescent",
	"opaque",
	"open",
	"opposite",
	"orange",
	"orchid",
	"orderly",
	"ordinary",
	"organic",
	"organized",
	"ossified",
	"outgoing",
	"outrageous",
	"outstanding",
	"oval",
	"overjoyed",
	"oxidized",
	"pacific",
	"paint",
	"painted",
	"pale",
	"palm",
	"panoramic",
	"paper",
	"parallel",
	"past",
	"pastoral",
	"patch",
	"pattern",
	"peaceful",
	"peach",
	"pear",
	"peat",
	"pebble",
	"pentagonal",
	"pepper",
	"peppered",
	"peppermint",
	"perfect",
	"peridot",
	"periodic",
	"periwinkle",
	"perpetual",
	"persistent",
	"petal",
	"petalite",
	"petite",
	"pewter",
	"phantom",
	"phase",
	"phrygian",
	"picayune",
	"pickle",
	"pickled",
	"picturesque",
	"pie",
	"pine",
	"pineapple",
	"pinnate",
	"pinto",
	"piquant",
	"pitch",
	"placid",
	"plaid",
	"plain",
	"planet",
	"plant",
	"plastic",
	"platinum",
	"plausible",
	"playful",
	"pleasant",
	"plucky",
	"plum",
	"plume",
	"plural",
	"pointed",
	"pointy",
	"poised",
	"polar",
	"polarized",
	"polished",
	"polite",
	"political",
	"pollen",
	"polydactyl",
	"polyester",
	"pond",
	"pool",
	"popular",
	"positive",
	"possible",
	"potent",
	"pouncing",
	"power",
	"powerful",
	"prairie",
	"precious",
	"pretty",
	"pricey",
	"prickle",
	"prickly",
	"principled",
	"prism",
	"private",
	"probable",
	"productive",
	"profuse",
	"prong",
	"protective",
	"proud",
	"proximal",
	"psychedelic",
	"puddle",
	"pumped",
	"purple",
	"purrfect",
	"purring",
	"pushy",
	"puzzle",
	"puzzled",
	"puzzling",
	"pyrite",
	"quaint",
	"quark",
	"quartz",
	"quasar",
	"quick",
	"quickest",
	"quiet",
	"quill",
	"quilled",
	"quilt",
	"quilted",
	"quintessential",
	"quirky",
	"quiver",
	"quixotic",
	"radial",
	"radical",
	"rain",
	"rainbow",
	"rainy",
	"rambunctious",
	"rapid",
	"rare",
	"raspy",
	"rattle",
	"real",
	"rebel",
	"recent",
	"receptive",
	"recondite",
	"rectangular",
	"reflective",
	"regal",
	"regular",
	"reinvented",
	"reliable",
	"relic",
	"relieved",
	"remarkable",
	"reminiscent",
	"repeated",
	"resilient",
	"resisted",
	"resolute",
	"resonant",
	"respected",
	"responsible",
	"rhetorical",
	"rhinestone",
	"ribbon",
	"rich",
	"rift",
	"right",
	"righteous",
	"rightful",
	"rigorous",
	"ring",
	"ringed",
	"ripe",
	"ripple",
	"ritzy",
	"river",
	"road",
	"roan",
	"roasted",
	"robust",
	"rocky",
	"rogue",
	"romantic",
	"roomy",
	"rose",
	"rotated",
	"rotating",
	"rough",
	"round",
	"rounded",
	"royal",
	"rumbling",
	"rune",
	"rural",
	"rust",
	"rustic",
	"saber",
	"sable",
	"safe",
	"sage",
	"salt",
	"salty",
	"same",
	"sand",
	"sandy",
	"sapphire",
	"sassy",
	"satin",
	"satisfying",
	"savory",
	"scalloped",
	"scandalous",
	"scarce",
	"scarlet",
	"scented",
	"scientific",
	"scintillating",
	"scratch",
	"scratched",
	"scrawny",
	"screeching",
	"scythe",
	"season",
	"seasoned",
	"second",
	"secret",
	"secretive",
	"sedate",
	"seed",
	"seemly",
	"seen",
	"selective",
	"separate",
	"separated",
	"sepia",
	"sequoia",
	"serene",
	"serious",
	"shade",
	"shaded",
	"shadow",
	"shadowed",
	"shared",
	"sharp",
	"sheer",
	"shell",
	"shelled",
	"shimmer",
	"shimmering",
	"shine",
	"shining",
	"shiny",
	"shocking",
	"shore",
	"short",
	"shorthaired",
	"showy",
	"shrouded",
	"shrub",
	"shy",
	"sideways",
	"silent",
	"silicon",
	"silk",
	"silken",
	"silky",
	"silly",
	"silver",
	"similar",
	"simple",
	"simplistic",
	"sincere",
	"single",
	"six",
	"sixth",
	"skillful",
	"skinny",
	"skitter",
	"sky",
	"slash",
	"sleepy",
	"sleet",
	"slender",
	"slime",
	"slimy",
	"slow",
	"sly",
	"small",
	"smart",
	"smiling",
	"smoggy",
	"smooth",
	"snapdragon",
	"sneaky",
	"snow",
	"snowy",
	"soapy",
	"soft",
	"solar",
	"solid",
	"solstice",
	"somber",
	"sophisticated",
	"sordid",
	"sore",
	"sour",
	"south",
	"southern",
	"spangle",
	"spangled",
	"spark",
	"sparkling",
	"sparkly",
	"special",
	"speckle",
	"speckled",
	"spectacled",
	"spectacular",
	"spectrum",
	"sphenoid",
	"spice",
	"spiced",
	"spicy",
	"spiffy",
	"spiky",
	"spiny",
	"spiral",
	"spiritual",
	"splashy",
	"splendid",
	"sponge",
	"spot",
	"spotless",
	"spotted",
	"spotty",
	"spring",
	"sprinkle",
	"sprout",
	"spurious",
	"square",
	"standing",
	"star",
	"statuesque",
	"steadfast",
	"steady",
	"stealth",
	"steel",
	"steep",
	"stellar",
	"sticky",
	"stingy",
	"stirring",
	"stitch",
	"stone",
	"storm",
	"stormy",
	"stream",
	"strengthened",
	"stripe",
	"striped",
	"strong",
	"stump",
	"stupendous",
	"sturdy",
	"stylish",
	"suave",
	"subdued",
	"subsequent",
	"substantial",
	"successful",
	"succinct",
	"succulent",
	"sudden",
	"sudsy",
	"sugar",
	"sugared",
	"sugary",
	"sulfuric",
	"sulky",
	"summer",
	"sumptuous",
	"sun",
	"sunny",
	"sunrise",
	"sunset",
	"super",
	"superb",
	"superficial",
	"supreme",
	"surf",
	"sustaining",
	"swamp",
	"swanky",
	"sweet",
	"sweltering",
	"swift",
	"synonymous",
	"tabby",
	"talented",
	"tall",
	"tame",
	"tan",
	"tangible",
	"tangy",
	"tar",
	"tarry",
	"tartan",
	"tasteful",
	"tasty",
	"tattered",
	"teal",
	"telling",
	"temporal",
	"ten",
	"tender",
	"terrific",
	"tested",
	"thankful",
	"therapeutic",
	"thin",
	"thinkable",
	"third",
	"thirsty",
	"thoracic",
	"thorn",
	"thoughtful",
	"thread",
	"three",
	"thrilling",
	"thunder",
	"thundering",
	"tidal",
	"tide",
	"tidy",
	"time",
	"tin",
	"tinted",
	"tiny",
	"titanium",
	"toothsome",
	"topaz",
	"torch",
	"torpid",
	"tortoiseshell",
	"tough",
	"tourmaline",
	"towering",
	"trail",
	"tranquil",
	"translucent",
	"transparent",
	"trapezoidal",
	"traveling",
	"treasure",
	"tree",
	"tremendous",
	"triangular",
	"tricky",
	"tricolor",
	"trite",
	"tropical",
	"troubled",
	"trusted",
	"trusting",
	"truth",
	"truthful",
	"tulip",
	"tundra",
	"tungsten",
	"turquoise",
	"twilight",
	"twisty",
	"typhoon",
	"typical",
	"ubiquitous",
	"ultra",
	"uncovered",
	"understood",
	"unequaled",
	"uneven",
	"unexpected",
	"unique",
	"universal",
	"unleashed",
	"unmarred",
	"unruly",
	"unusual",
	"upbeat",
	"useful",
	"utopian",
	"uttermost",
	"vagabond",
	"valiant",
	"valley",
	"valuable",
	"vanilla",
	"various",
	"vast",
	"vaulted",
	"veil",
	"veiled",
	"verbena",
	"verbose",
	"verdant",
	"versed",
	"victorious",
	"vigorous",
	"vine",
	"vintage",
	"violet",
	"viridian",
	"visual",
	"vivacious",
	"vivid",
	"volcano",
	"voltaic",
	"voracious",
	"waiting",
	"wakeful",
	"walnut",
	"wandering",
	"warm",
	"warp",
	"wary",
	"water",
	"watery",
	"wave",
	"wax",
	"wealthy",
	"well",
	"west",
	"western",
	"wheat",
	"whimsical",
	"whip",
	"whispering",
	"wholesale",
	"wide",
	"wiggly",
	"wild",
	"wind",
	"winter",
	"wirehaired",
	"wiry",
	"wise",
	"wistful",
	"witty",
	"wobbly",
	"wonderful",
	"wood",
	"wooded",
	"wooden",
	"wool",
	"woolen",
	"woolly",
	"workable",
	"working",
	"worried",
	"wry",
	"yellow",
	"yielding",
	"young",
	"youthful",
	"zany",
	"zealous",
	"zenith",
	"zest",
	"zesty",
	"zigzag",
	"zinc",
	"zippy",
	"zircon",
}

var WORDLIST_PREDICATES_LEN = len(WORDLIST_PREDICATES)