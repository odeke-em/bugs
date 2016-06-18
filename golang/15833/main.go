package main

import (
	"flag"
	"log"
	"math/big"
	"net/http"
	"runtime/trace"
	"time"
)

var (
	intZero = big.NewInt(0)
	intOne  = big.NewInt(1)
)

func AlgoGCD(a, b *big.Int) *big.Int {
	aZeroCmp := a.Cmp(intZero)
	if aZeroCmp <= 0 {
		return b
	}
	bZeroCmp := b.Cmp(intZero)
	if bZeroCmp <= 0 {
		return a
	}

	aCmp := a.Cmp(b)
	switch aCmp {
	case 0:
		return a
	case 1:
		// Swap a and b
		a, b = b, a
	}

	closestWhole := big.NewInt(0).Div(b, a)

	bn := big.NewInt(0)
	an := big.NewInt(0)
	residue := bn.Sub(b, an.Mul(a, closestWhole))

	if residue.Cmp(intZero) == 0 {
		return a
	}

	return AlgoGCD(a, residue)
}

func traceResult(w http.ResponseWriter, r *http.Request) {
	trace.Start(w)
	defer trace.Stop()

	a, _ := new(big.Int).SetString("162446598402363997673493179860749206481570626148774722480464916217519998582106439959196344215073041683672767407342600465416394761370312812652387014077287533579464234592025270747189685651849855568219894257068017511394846018248946392555048664581847849090691446151175543365199139365693828091221296651564817601708173519333480541557916813321048094446738924484963396047875101501912007668905398696596051247666259795425384659975708544211774775625965457167714100319140254438470405084497675188274084057803534621424329554676703869536840619010332513253366351288042987043711316273802950111102675169897156927510722936205243934593761221732977646338069405205393572558504296440846844654915994056051080138684289188618799496532006547298248428910424360291035687113795912784248641162116970198942033380494094025012100218442198181604277428406480729238384656494840163546922286562841336237622223094660753989052636783736736730051815788481272143711008904458168016828251588196506829619252076432097202449298271511577883563293048604787739890118213783329525568780738050849374005547963863874473115226603810373230563816359760094564252571054145767564049725284205901365726049711717022372349529374435379628254191205730293550748181543891163723281674777750059512413586588828314084535671582111735660379159552799083036587458583555292293962217151577840489845885188592504493431102401049590622433757873291425858434918697762180295687147582650879514071801716939533681208102052410202696047460966976913845601990028270324422645505289049580017114480330931133826162290356916061574256540132190171694393638181573510025576808968516762515657212669911734833269039513284253725561578543014863535935809511150296456155470508927211411840144748006797665686217130525378029137093695488159569875226988769985175538075995372140625158000424680862039472332990803861888330870444263274715912499029641167624951512824402013661747715825752288086558167740183961939884790195081886224217177152869350897055478137773963535642498587936119572587325769768711699791252679336168006908091872154677467014758193802812900921873466740528890306650385439501239427958166273861080306246501556064395044538359837700353874621248486133241226805310121486599455004347652284602607212474583483132546550311381260917418037865509077618272877792484328607271857806356", 10)
	b, _ := new(big.Int).SetString("101784352192926179474691847690295364412146644944337147893018137607751766917729750758328166579925569899765459590972517709626629422864683438193290125043063518436017102066763671846630815217722251657259307553099132772201209113951039569557515190691205890052686828351751884803712581582093817758343808197167567247907540122605321423275330224580551817963867742145560163431380567362003538753856228060374866951036259004102917220472306669995666505172223628113212045865085784184336530203837908112187832616046195538504845768119298751526550607442806456808908879070037709190448307774015887964660855587392536795762810554496174847323701119347378258018659963334626467144614011455875192800289417692534392258372638105441472555031300945385673635743106163778851484848839379940456113937207911595141513680574367846648636353220309193154584372103496857062568070324916862436127133316127989083231670428229702525886963156897871633715273896484904539939624236288247009844489720556858425695244949730267384132132441488654009449316215449359564589915898961035516198007811431244695623908883885142048155515509053435754497981693091320819719426795331186500179253864761247316755487225443097006915430516922716275267385443349373289960634713632572446080364203373228601555282278183545440347991778185107399327758589346260248433939637585528206559787605105051710971698307363677567320558959899872756874248500394667863120050131533976493034552093698090232503340912301095656404952421422358796326151351099379863658968843482129929673487863396575147218726229000469299309801093105964119662371898258712567149943154476612246157910925561894764017985426180288097588055722886924974722109733733123413154279708050676188619796030777587411061473509487256419260726999082015872941209426850601596056518119145537343224544999826199864836904118369929404501365516712854950664599501540383745231341313574626357265253728062815992109070862862165578453391683605892693204046869207638687457235701255574852208310767024525571635865352089126234994777689748203394347241246694312344558416290642852901684770713683262440625612790798884009774609112528636374285113491984040948539722165924050855717036844286332805554836911465419002994297397099377669991903562736772713728306241921475535135050281780162431545636315231851527333384539192469227100575974540", 10)

	origGCD := false
	n := 1
	flag.BoolVar(&origGCD, "use-orig", false, "use the original big GCD algo")
	flag.IntVar(&n, "n", 10000, "n iterations")
	flag.Parse()

	startTime := time.Now()

	if false {
		a = big.NewInt(1096 * 1096)
		b = big.NewInt(51 * 2918273)
	}

	for i := 0; i < n; i++ {
		new(big.Int).GCD(nil, nil, a, b)
	}

	log.Println(time.Since(startTime))
}

func main() {
	server := new(http.Server)
	server.Addr = ":8888"
	http.HandleFunc("/", traceResult)
	log.Printf("running server on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
