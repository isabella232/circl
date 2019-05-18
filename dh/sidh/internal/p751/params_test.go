package p751

// Contains values used by tests
import (
	"testing/quick"

	. "github.com/cloudflare/circl/dh/sidh/internal/common"
)

var (
	expectedJ = Fp2{
		A: Fp{0xc7a8921c1fb23993, 0xa20aea321327620b, 0xf1caa17ed9676fa8, 0x61b780e6b1a04037, 0x47784af4c24acc7a, 0x83926e2e300b9adf, 0xcd891d56fae5b66, 0x49b66985beb733bc, 0xd4bcd2a473d518f, 0xe242239991abe224, 0xa8af5b20f98672f8, 0x139e4d4e4d98},
		B: Fp{0xb5b52a21f81f359, 0x715e3a865db6d920, 0x9bac2f9d8911978b, 0xef14acd8ac4c1e3d, 0xe81aacd90cfb09c8, 0xaf898288de4a09d9, 0xb85a7fb88c5c4601, 0x2c37c3f1dd303387, 0x7ad3277fe332367c, 0xd4cbee7f25a8e6f8, 0x36eacbe979eaeffa, 0x59eb5a13ac33},
	}
	// A = 4385300808024233870220415655826946795549183378139271271040522089756750951667981765872679172832050962894122367066234419550072004266298327417513857609747116903999863022476533671840646615759860564818837299058134292387429068536219*i + 1408083354499944307008104531475821995920666351413327060806684084512082259107262519686546161682384352696826343970108773343853651664489352092568012759783386151707999371397181344707721407830640876552312524779901115054295865393760
	curve_A = Fp2{
		A: Fp{0x8319eb18ca2c435e, 0x3a93beae72cd0267, 0x5e465e1f72fd5a84, 0x8617fa4150aa7272, 0x887da24799d62a13, 0xb079b31b3c7667fe, 0xc4661b150fa14f2e, 0xd4d2b2967bc6efd6, 0x854215a8b7239003, 0x61c5302ccba656c2, 0xf93194a27d6f97a2, 0x1ed9532bca75},
		B: Fp{0xb6f541040e8c7db6, 0x99403e7365342e15, 0x457e9cee7c29cced, 0x8ece72dc073b1d67, 0x6e73cef17ad28d28, 0x7aed836ca317472, 0x89e1de9454263b54, 0x745329277aa0071b, 0xf623dfc73bc86b9b, 0xb8e3c1d8a9245882, 0x6ad0b3d317770bec, 0x5b406e8d502b}}

	// C = 933177602672972392833143808100058748100491911694554386487433154761658932801917030685312352302083870852688835968069519091048283111836766101703759957146191882367397129269726925521881467635358356591977198680477382414690421049768*i + 9088894745865170214288643088620446862479558967886622582768682946704447519087179261631044546285104919696820250567182021319063155067584445633834024992188567423889559216759336548208016316396859149888322907914724065641454773776307
	curve_C = Fp2{
		A: Fp{0x4fb2358bbf723107, 0x3a791521ac79e240, 0x283e24ef7c4c922f, 0xc89baa1205e33cc, 0x3031be81cff6fee1, 0xaf7a494a2f6a95c4, 0x248d251eaac83a1d, 0xc122fca1e2550c88, 0xbc0451b11b6cfd3d, 0x9c0a114ab046222c, 0x43b957b32f21f6ea, 0x5b9c87fa61de},
		B: Fp{0xacf142afaac15ec6, 0xfd1322a504a071d5, 0x56bb205e10f6c5c6, 0xe204d2849a97b9bd, 0x40b0122202fe7f2e, 0xecf72c6fafacf2cb, 0x45dfc681f869f60a, 0x11814c9aff4af66c, 0x9278b0c4eea54fe7, 0x9a633d5baf7f2e2e, 0x69a329e6f1a05112, 0x1d874ace23e4}}

	// x(P) = 8172151271761071554796221948801462094972242987811852753144865524899433583596839357223411088919388342364651632180452081960511516040935428737829624206426287774255114241789158000915683252363913079335550843837650671094705509470594*i + 9326574858039944121604015439381720195556183422719505497448541073272720545047742235526963773359004021838961919129020087515274115525812121436661025030481584576474033630899768377131534320053412545346268645085054880212827284581557
	affine_xP = Fp2{
		A: Fp{0xe8d05f30aac47247, 0x576ec00c55441de7, 0xbf1a8ec5fe558518, 0xd77cb17f77515881, 0x8e9852837ee73ec4, 0x8159634ad4f44a6b, 0x2e4eb5533a798c5, 0x9be8c4354d5bc849, 0xf47dc61806496b84, 0x25d0e130295120e0, 0xdbef54095f8139e3, 0x5a724f20862c},
		B: Fp{0x3ca30d7623602e30, 0xfb281eddf45f07b7, 0xd2bf62d5901a45bc, 0xc67c9baf86306dd2, 0x4e2bd93093f538ca, 0xcfd92075c25b9cbe, 0xceafe9a3095bcbab, 0x7d928ad380c85414, 0x37c5f38b2afdc095, 0x75325899a7b779f4, 0xf130568249f20fdd, 0x178f264767d1}}

	// x(Q) = 613162677562606602867371958793876971530136728660199185642812914351735322828980793298955764087721218995329689800176835576008483782305021671417330230966*i + 12939479853552958669415184882894789433224467828463058020095712989798268661016843892597050485554155971441665589419365957826417334087966365414056706471602
	affine_xQ = Fp2{
		A: Fp{0xd3d14533cb0db45c, 0xdaf10b9f5fb037cf, 0x9562c31985823562, 0xb79b75e2bf5635a5, 0x83d38fb1669c2d9, 0x5e67117a35a57cbd, 0x4b94ed6c3cbf54a4, 0x3f47706b62778d},
		B: Fp{0x1c766c0e7ed612d6, 0x2f3b42979e8efd86, 0xd82bac0006415ee7, 0x20cfe3bec0970042, 0x8f6628807e862bf9, 0xac4f82d13fddd9c5, 0x70b756e4bac1fa85, 0x350c02508e50dc}}

	affine_xPmQ = Fp2{
		A: Fp{0x7295ac0e8a5531c8, 0xc02afbef6cdf51a8, 0x84d0c1bb8d80624f, 0x26abbf06c61e814b, 0xada277883fbdae07, 0x63b6739db68df3a7, 0x3554670381bcfcc6, 0x33419be255bbc6},
		B: Fp{0xff9df35102da997a, 0xf9b920b2d8bd6210, 0x43d1514920e73bfb, 0x624c6fd6ef16da74, 0xeb535968ed0d1286, 0x9243e6ce29a636e6, 0x1849ed96cb7940e0, 0x1e4b495933b675}}

	// x([2]P) = 1476586462090705633631615225226507185986710728845281579274759750260315746890216330325246185232948298241128541272709769576682305216876843626191069809810990267291824247158062860010264352034514805065784938198193493333201179504845*i + 3623708673253635214546781153561465284135688791018117615357700171724097420944592557655719832228709144190233454198555848137097153934561706150196041331832421059972652530564323645509890008896574678228045006354394485640545367112224
	affine_xP2 = Fp2{
		A: Fp{0x2a77afa8576ce979, 0xab1360e69b0aeba0, 0xd79e3e3cbffad660, 0x5fd0175aa10f106b, 0x1800ebafce9fbdbc, 0x228fc9142bdd6166, 0x867cf907314e34c3, 0xa58d18c94c13c31c, 0x699a5bc78b11499f, 0xa29fc29a01f7ccf1, 0x6c69c0c5347eebce, 0x38ecee0cc57},
		B: Fp{0x43607fd5f4837da0, 0x560bad4ce27f8f4a, 0x2164927f8495b4dd, 0x621103fdb831a997, 0xad740c4eea7db2db, 0x2cde0442205096cd, 0x2af51a70ede8324e, 0x41a4e680b9f3466, 0x5481f74660b8f476, 0xfcb2f3e656ff4d18, 0x42e3ce0837171acc, 0x44238c30530c}}

	// x([3]P) = 9351941061182433396254169746041546943662317734130813745868897924918150043217746763025923323891372857734564353401396667570940585840576256269386471444236630417779544535291208627646172485976486155620044292287052393847140181703665*i + 9010417309438761934687053906541862978676948345305618417255296028956221117900864204687119686555681136336037659036201780543527957809743092793196559099050594959988453765829339642265399496041485088089691808244290286521100323250273
	affine_xP3 = Fp2{
		A: Fp{0x2096e3f23feca947, 0xf36f635aa4ad8634, 0xdae3b1c6983c5e9a, 0xe08df6c262cb74b4, 0xd2ca4edc37452d3d, 0xfb5f3fe42f500c79, 0x73740aa3abc2b21f, 0xd535fd869f914cca, 0x4a558466823fb67f, 0x3e50a7a0e3bfc715, 0xf43c6da9183a132f, 0x61aca1e1b8b9},
		B: Fp{0x1e54ec26ea5077bd, 0x61380572d8769f9a, 0xc615170684f59818, 0x6309c3b93e84ef6e, 0x33c74b1318c3fcd0, 0xfe8d7956835afb14, 0x2d5a7b55423c1ecc, 0x869db67edfafea68, 0x1292632394f0a628, 0x10bba48225bfd141, 0x6466c28b408daba, 0x63cacfdb7c43}}

	// x([2^2]P) = 441719501189485559222919502512761433931671682884872259563221427434901842337947564993718830905758163254463901652874331063768876314142359813382575876106725244985607032091781306919778265250690045578695338669105227100119314831452*i + 6961734028200975729170216310486458180126343885294922940439352055937945948015840788921225114530454649744697857047401608073256634790353321931728699534700109268264491160589480994022419317695690866764726967221310990488404411684053
	affine_xP4 = Fp2{
		A: Fp{0x6f9dbe4c39175153, 0xf2fec757eb99e88, 0x43d7361a93733d91, 0x3abd10ed19c85a3d, 0xc4de9ab9c5ef7181, 0x53e375901684c900, 0x68ffc3e7d71c41ff, 0x47adab62c8d942fe, 0x226a33fd6fbb381d, 0x87ef4c8fdd83309a, 0xaca1cf44c5fa8799, 0x6cbae86c755f},
		B: Fp{0x4c80c37fe68282a7, 0xbd8b9d7248bf553a, 0x1fb0e8e74d5e1762, 0xb63fa0e4e5f91482, 0xc675ab8a45a1439, 0xdfa6772deace7820, 0xf0d813d71d9a9255, 0x53a1a58c634534bd, 0x4ebfc6485fdfd888, 0x6991fe4358bcf169, 0xc0547bdaca85b6fd, 0xf461548d632}}

	// x([3^2]P) = 3957171963425208493644602380039721164492341594850197356580248639045894821895524981729970650520936632013218950972842867220898274664982599375786979902471523505057611521217523103474682939638645404445093536997296151472632038973463*i + 1357869545269286021642168835877253886774707209614159162748874474269328421720121175566245719916322684751967981171882659798149072149161259103020057556362998810229937432814792024248155991141511691087135859252304684633946087474060
	affine_xP9 = Fp2{
		A: Fp{0x7c0daa0f04ded4e0, 0x52dc4f883d85e065, 0x91afbdc2c1714d0b, 0xb7b3db8e658cfeba, 0x43d4e72a692882f3, 0x535c56d83753da30, 0xc8a58724433cbf5d, 0x351153c0a5e74219, 0x2c81827d19f93dd5, 0x26ef8aca3370ea1a, 0x1cf939a6dd225dec, 0x3403cb28ad41},
		B: Fp{0x93e7bc373a9ff7b, 0x57b8cc47635ebc0f, 0x92eab55689106cf3, 0x93643111d421f24c, 0x1c58b519506f6b7a, 0xebd409fb998faa13, 0x5c86ed799d09d80e, 0xd9a1d764d6363562, 0xf95e87f92fb0c4cc, 0x6b2bbaf5632a5609, 0x2d9b6a809dfaff7f, 0x29c0460348b}}

	// Inputs for testing 3-point-ladder
	threePointLadderInputs = []ProjectivePoint{
		// x(P)
		ProjectivePoint{
			X: Fp2{
				A: Fp{0xe8d05f30aac47247, 0x576ec00c55441de7, 0xbf1a8ec5fe558518, 0xd77cb17f77515881, 0x8e9852837ee73ec4, 0x8159634ad4f44a6b, 0x2e4eb5533a798c5, 0x9be8c4354d5bc849, 0xf47dc61806496b84, 0x25d0e130295120e0, 0xdbef54095f8139e3, 0x5a724f20862c},
				B: Fp{0x3ca30d7623602e30, 0xfb281eddf45f07b7, 0xd2bf62d5901a45bc, 0xc67c9baf86306dd2, 0x4e2bd93093f538ca, 0xcfd92075c25b9cbe, 0xceafe9a3095bcbab, 0x7d928ad380c85414, 0x37c5f38b2afdc095, 0x75325899a7b779f4, 0xf130568249f20fdd, 0x178f264767d1}},
			Z: params.OneFp2,
		},
		// x(Q)
		ProjectivePoint{
			X: Fp2{
				A: Fp{0x2b71a2a93ad1e10e, 0xf0b9842a92cfb333, 0xae17373615a27f5c, 0x3039239f428330c4, 0xa0c4b735ed7dcf98, 0x6e359771ddf6af6a, 0xe986e4cac4584651, 0x8233a2b622d5518, 0xbfd67bf5f06b818b, 0xdffe38d0f5b966a6, 0xa86b36a3272ee00a, 0x193e2ea4f68f},
				B: Fp{0x5a0f396459d9d998, 0x479f42250b1b7dda, 0x4016b57e2a15bf75, 0xc59f915203fa3749, 0xd5f90257399cf8da, 0x1fb2dadfd86dcef4, 0x600f20e6429021dc, 0x17e347d380c57581, 0xc1b0d5fa8fe3e440, 0xbcf035330ac20e8, 0x50c2eb5f6a4f03e6, 0x86b7c4571}},
			Z: params.OneFp2,
		},
		// x(P-Q)
		ProjectivePoint{
			X: Fp2{
				A: Fp{0x4aafa9f378f7b5ff, 0x1172a683aa8eee0, 0xea518d8cbec2c1de, 0xe191bcbb63674557, 0x97bc19637b259011, 0xdbeae5c9f4a2e454, 0x78f64d1b72a42f95, 0xe71cb4ea7e181e54, 0xe4169d4c48543994, 0x6198c2286a98730f, 0xd21d675bbab1afa5, 0x2e7269fce391},
				B: Fp{0x23355783ce1d0450, 0x683164cf4ce3d93f, 0xae6d1c4d25970fd8, 0x7807007fb80b48cf, 0xa005a62ec2bbb8a2, 0x6b5649bd016004cb, 0xbb1a13fa1330176b, 0xbf38e51087660461, 0xe577fddc5dd7b930, 0x5f38116f56947cd3, 0x3124f30b98c36fde, 0x4ca9b6e6db37}},
			Z: params.OneFp2,
		},
	}
	scalar3Pt = [...]uint8{84, 222, 146, 63, 85, 18, 173, 162, 167, 38, 10, 8, 143, 176, 93, 228, 247, 128, 50, 128, 205, 42, 15, 137, 119, 67, 43, 3, 61, 91, 237, 24, 235, 12, 53, 96, 186, 164, 232, 223, 197, 224, 64, 109, 137, 63, 246, 4}
)

var quickCheckConfig = &quick.Config{
	MaxCount: (1 << 15),
}
