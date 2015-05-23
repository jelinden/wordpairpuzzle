###Get the source code
```
go get github.com/jelinden/wordpairpuzzle
```

###Build
```
go build
```

###Run
```
GOMAXPROCS=3 ./wordpairpuzzle
```
GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and 
returns the previous setting. If n < 1, it does not change the current setting.







##KOODAUSPÄHKINÄ, KESÄ 2015: MUHKEIMMAT SANAPARIT

Tällä kertaa tutkimme muhkeita sanapareja.

Sanaparin muhkeus määritellään seuraavasti: muhkeus on yhtä kuin sanaparin sisältämien uniikkien kirjainten määrä. 
Kirjaimiksi lasketaan seuraavat merkit: {a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, å, ä, ö}. 
Isot ja pienet kirjaimet lasketaan samaksi. Esimerkiksi sanaparin {"Upea", "Kapteeni"} muhkeus on 8, 
koska sanapari sisältää kirjaimet {a, e, i, k, n, p, t, u}.

Kysymys kuuluu: mikä on Alastalon salissa -kirjan muhkein sanapari, tai muhkeimmat sanaparit, 
jos useampi pari saa korkeimmat muhkeuspisteet? 
Sanojen ei tarvitse olla peräkkäin, vaan tarkoitus on löytää koko kirjan kaikista sanoista ne kaksi sanaa, 
jotka muodostavat muhkeimman parin.

##VINKKEJÄ

Tämänkertainen pähkinä on paljon vaikeampi kuin edellinen. Puhtaasti brute forcella tämä ei onnistu.

##SÄÄNNÖT

Osallistua saa millä ohjelmointikielellä tahansa.

Voittajaksi valitaan elegantein ratkaisu. Voittajaratkaisulta vaaditaan eleganttiuden lisäksi, 
että se toimii oikein ja antaa oikean vastauksen.