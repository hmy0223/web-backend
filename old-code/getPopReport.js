import { sleep } from 'k6';
import { parseHTML } from 'k6/html';
import http from 'k6/http';

//const cardData = JSON.parse(open("./temp.json"));

//let vu = cardData.cardNumber.length
let vu = 50

export let options = {
    scenarios: {
    //   loading: {
    //     executor: 'per-vu-iterations',
    //     maxDuration: "1m",
    //     vus: vu,
    //     iterations: 1,
    //   }
    // }

    constant_request_rate: {
        executor: "constant-arrival-rate",
  
        rate: 7000,
        timeUnit: "1.3",  
        preAllocatedVUs: 1,
        maxVUs: 1,
  
        duration: "210m"

      },
    },
};
    
export default function () {
    // let jsonIndex = __ITER
    // let cardSerialNumberBase = cardData.cardNumber[jsonIndex].cardSerialNumber
    // let cardSerialNumber = cardSerialNumberBase.toString()
    sleep(0.5)
    let cardSerialNumberBase = 15310680 
	let cardSerialNumberLoop = cardSerialNumberBase + (__ITER + 1)
    let cardSerialNumber = cardSerialNumberLoop.toString()
    
    let cardLookUpUrl = `https://www.beckett.com/grading/card-lookup?item_type=BGS&item_id=${cardSerialNumber}&submit=Submit&submit=Submit`;

    let res = http.get(cardLookUpUrl);
    let doc = parseHTML(res.body);

    let setName = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(1) > td:nth-child(3)').text();
    let cardName = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(2) > td:nth-child(3)').text();
    let dateGraded = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(3) > td:nth-child(3)').text()
    
    let centeringGrade = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(4) > td:nth-child(3)').text()
    let cornerGrade = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(5) > td:nth-child(3)').text()
    let edgesGrade = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(6) > td:nth-child(3)').text()
    let surfaceGrade = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(7) > td:nth-child(3)').text()
    
    let finalGrade = doc.find('#grading > div > div > div.main_content_area > table > tbody > tr:nth-child(9) > td:nth-child(3)').text()

    if(centeringGrade == '10.0' && cornerGrade == '10.0' && edgesGrade == '10.0' && surfaceGrade == '10.0') {
        finalGrade = '10.1'
    }

    let cardDetails = {
        "Card Serial Number": cardSerialNumber,
        "Set Name": setName,
        "Card Name": cardName,
        "Date Graded": dateGraded,
        "Grade": finalGrade
    }

    if(cardDetails['Set Name'] != ""){
        console.log(cardDetails)
    }
    // if(cardDetails['Set Name'] == "2022 Pokemon CoroCoro Start Deck 100"){
    //     console.log(cardDetails)
    // }
}