package main

import (
	"encoding/json"

	"html/template"
	"io/ioutil"
	//"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
	//"cloud.google.com/go/storage"
	//"golang.org/x/net/context"
	//"google.golang.org/api/option"
)

//funzione di errore
func checkErrors(err error) {
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

//ReadFromJSON function load a json file into a struct or return error
func ReadFromJSON(t interface{}, filename string) error {

	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonFile), t)
	if err != nil {
		return err
	}

	return nil
}

//Crea una struttura dal json
type Datas struct {
	Classe          []string     `json:"classe"`
	Genere          []string     `json:"genere"`
	Razza           []string     `json:"razza"`
	Allineamento    []string     `json:"allineamento"`
	Taglia          []string     `json:"taglia"`
	Dio             []string     `json:"dio"`
	NomePersonaggio [][][]string `json:"nomePersonaggio"`
}

var Conf Datas

//genera scheda del personaggio randomizzata
func Genera(pm map[string]string, s []int) map[string]string {

	selezioneNome := Conf.NomePersonaggio[s[0]][s[1]] // crea il nome del personaggio basandosi su razza e genere
	rand.Seed(time.Now().UnixNano())

	pm["Nome"] = selezioneNome[rand.Intn(len(selezioneNome))]
	pm["Allineamento"] = Conf.Allineamento[rand.Intn(len(Conf.Allineamento))]
	pm["Taglia"] = Conf.Taglia[rand.Intn(len(Conf.Taglia))]
	pm["Classe"] = Conf.Classe[rand.Intn(len(Conf.Classe))]
	pm["Divinità"] = Conf.Dio[rand.Intn(len(Conf.Dio))]

	return pm
}

/*func pdfCreate(pdm map[string]string) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "") //crea il pdf
	pdf.AddPage()                          //crea la pagina
	pdf.SetFont("Arial", "B", 12)          //imposta il font

	pdf.Cell(40, 10, "Giocatore")   //crea il Nome Giocatore
	pdf.Cell(40, 10, "Personaggio") //crea il Nome Personaggio
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, pdm["Utente"]) // rimpie il nome giocatore
	pdf.Cellf(40, 10, pdm["Nome"])  //	riempie il nome personaggio
	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 12) //imposta il font
	pdf.Cell(40, 10, "Razza")     //crea la razza
	pdf.Cell(40, 10, "Genere")    //crea il genere
	pdf.Ln(8)                     //a capo (spaziatura normale)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, pdm["Razza"])   // rimpie il la razza
	pdf.Cell(40, 10, pdm["Genere"])  // rimpie il genere
	pdf.Ln(8)                        //a capo (spaziatura normale)
	pdf.SetFont("Arial", "B", 12)    //imposta il font
	pdf.Cell(40, 10, "Allineamento") //crea l'allineamento
	pdf.Cell(40, 10, "Taglia")       //crea la taglia
	pdf.Cell(40, 10, "Classe")       //crea la classe
	pdf.Cell(40, 10, "Divinita'")    //crea la divinità
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, pdm["Allineamento"]) // rimpie l'allineamento
	pdf.Cell(40, 10, pdm["Taglia"])       // rimpie la taglia
	pdf.Cell(40, 10, pdm["Classe"])       // rimpie la classe
	pdf.Cell(40, 10, pdm["Divinità"])     // rimpie la divinità
	//checkErrors(pdf.OutputFileAndClose(pdm["Utente"] + ".pdf")) //salva il pdf
	pdf.Close()

	return pdf
}*/

/*
func bucketSave(){

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	bkt := client.Bucket("dndrandomsheet.appspot.com")

	obj := bkt.Object("data")
	// Write something to obj.
	// w implements io.Writer.
	log.Println("numero1")
	w := obj.NewWriter(ctx)
	// Write some text to obj. This will either create the object or overwrite whatever is there already.
	if _, err := fmt.Fprintf(w, "This object contains text.\n"); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Println("numero2")
	// Close, just like writing a file.
	if err := w.Close(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Println("numero3")
	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Println("numero4")
	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Prints "This object contains text."
	acls, err := obj.ACL().List(ctx)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	for _, rule := range acls {
		fmt.Printf("%s has role %s\n", rule.Entity, rule.Role)
	}

}
*/
//legge il json
func init() {
	checkErrors(ReadFromJSON(&Conf, "conf.json"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("dnd.html"))
	homeMap := make(map[string]interface{})
	homeMap["Razza"] = Conf.Razza
	homeMap["Genere"] = Conf.Genere
	checkErrors(tmpl1.Execute(w, homeMap))

}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	tmp2 := template.Must(template.ParseFiles("answer.html"))
	processMap := make(map[string]string) //mappa per salvare i parametri

	processMap["Utente"] = r.FormValue("firstname")
	convertiGenere, _ := strconv.Atoi(r.FormValue("genere"))
	convertiRazza, _ := strconv.Atoi(r.FormValue("razza"))

	selezioni := []int{convertiGenere, convertiRazza}
	processMap["Genere"] = Conf.Genere[convertiGenere]
	processMap["Razza"] = Conf.Razza[convertiRazza]
	Genera(processMap, selezioni)
	//name :=processMap["Utente"] + ".pdf"
	//pdfCreate(processMap)
	//bucketSave()
	//http.ServeFile(w, r, name)
	//tmp2.Execute(w,pda)
	//downloadBytes,_ :=ioutil.ReadFile(name)
	//mime := http.DetectContentType(downloadBytes)

	//checkErrors(pda.Output(w))

	//fileSize := len(string(downloadBytes))
	//w.Header().Set("Content-Type", mime)
	//w.Header().Set("Content-Disposition", "attachment; filename="+name+"")
	//w.Header().Set("Expires", "0")
	//w.Header().Set("Content-Transfer-Encoding", "binary")
	//w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	//log.Println(fileSize)
	//http.ServeContent(w, r, name, time.Now(), bytes.NewReader(downloadBytes))

	checkErrors(tmp2.Execute(w, processMap))
}

func main() {

	http.HandleFunc("/", homeHandler)            //handler della pagina home
	http.HandleFunc("/process", answerHandler)   //handler della pagina di risposta /process
	log.Fatal(http.ListenAndServe(":8080", nil)) //hosting pagina
}
