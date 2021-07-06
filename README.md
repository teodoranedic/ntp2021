# Projekat iz predmeta Napredne tehnike programiranja

## Problem n-tela
Sekvencijalna i paralelizovana verzija rešenja u Python-u i Go-u.

### Opis problema
Simulacija n-tela aproksimira kretanje čestica koje su u međusobnoj interakciji putem fizičkih sila. Na osnovu ovako široke definicije, te čestice mogu biti u rasponu od nebeskih tela do atoma u oblaku gasa. Specijalizacija problema n-tela koji je rešavan u ovom projektu je vezana za gravitacione interakcije nebeskih tela u Sunčevom sistemu (n je 10, zajedno sa Suncem i Plutonom kao planetom patuljkom).  

### Algoritam - brute force
Problem se rešava upotrebom Njutnove jednačine gravitacione sile, gde se ukupna sila koja utiče na svaku česticu računa sumiranjem sila koje daju pojedinačne čestice u sistemu. Računanjem ukupne sile za svaku česticu, može se dalje dobiti brzina i pozicija čestice, korišćenje diskretizovanog vremenskog koraka (dt). Ova metoda je brute-force, jedina aproksimacija je da se ubrzanje čestica smatra konstantnim tokom vremenskog koraka. Ukoliko je korak dovoljno mali, aproksimacija postaje validna, ali to ima cenu većeg broja izračunavanja. 
Glavni nedostatak ovog algoritma je asimptotsko vreme računanja koje iznosi N<sup>2</sup> gde je N broj čestica, što znači da ukoliko dupliramo broj čestica, vreme računanja će biti 4 puta veće. Za veliki broj čestica ovaj algoritam je previše spor, pa se koriste druge metode kao što je Barnes-Hut, rekurzivni algoritam koji je zasnovan na kreiranju posebne strukture podataka oct-tree.

### Numeričke metode
Numeričke metode koje su primenjene za računanje brzina i pozicija čestica su Ojlerova i Leapfrog (žablji skok) metoda. Za inicijalni korak se koristi Ojlerova metoda za procenu brzine i pozicije (jer nemamo i-1 vrednosti koje su potrebne za Leapfrog):

  * v<sub>i+1</sub> = v<sub>i</sub>+a<sub>i</sub> dt 
  * p<sub>i+1</sub> = p<sub>i</sub>+v<sub>i</sub> dt 
  
U daljim iteracijama se koristi Leapfrog metoda:
  * v<sub>i+1</sub> = v<sub>i-1</sub>+2 a<sub>i</sub> dt
  * p<sub>i+1</sub> = p<sub>i-1</sub>+2 v<sub>i</sub> dt  
  
Na osnovu novih pozicija se u svakoj iteraciji računa ubrzanje pomoću Njutnovog zakona gravitacije:
  * F = G m<sub>1</sub> m<sub>2</sub>/r<sup>2</sup>
gde je G gravitaciona konstanta, m1 i m2 mase čestica i r rastojanje izmedju njih.

### Paralelizacija
Paralelizacija u Python jeziku je rađena pomoću multiprocessing paketa.

### Vizualizacija rešenja
Vizualizacija je zamišljena kao sekvenca slika nakon svake iteracije.

## Student
Teodora Nedić, sw41/2017
