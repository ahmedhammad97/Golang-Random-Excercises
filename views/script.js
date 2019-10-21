const questions = document.querySelector(".questions");
const backBtn = document.querySelector("#back");

document.querySelectorAll(".btn").forEach(el => {
  el.addEventListener('click', e => {
    questions.style.display = "none";
    backBtn.style.display = "block";
    document.querySelector("." + el.id).style.display = "block";
  });
});

backBtn.addEventListener('click', e => {
  document.querySelectorAll(".question_div").forEach(div => {
    div.style.display = "none";
  })
  questions.style.display = "block";
  backBtn.style.display = "none";
})

document.querySelector("#leap_submit").addEventListener('click', e => {
  var year = document.querySelector(".leap_div input").value;
  sendRequest("/leap", {"year" : year}).then(result => {
    document.querySelector(".leap_div .result").innerText = result;
  }).catch(err => {console.log(err)});
})

document.querySelector("#lived_submit").addEventListener('click', e => {
  var time = document.querySelector(".lived_div input").value + ":04Z";
  sendRequest("/lived", {"time" : time}).then(result => {
    document.querySelector(".lived_div .result").innerText = result.slice(0, result.length - 9);
  }).catch(err => {console.log(err)});
})

document.querySelector("#squares_submit").addEventListener('click', e => {
  var num = document.querySelector(".squares_div input").value;
  sendRequest("/squares", {"num" : num}).then(result => {
    document.querySelector(".squares_div .result").innerText = result;
  }).catch(err => {console.log(err)});
})

document.querySelector("#multiples_submit").addEventListener('click', e => {
  var nums = document.querySelectorAll(".multiples_div input");
  sendRequest("/multiples", {"num" : nums[0].value, "N" : nums[1].value}).then(result => {
    document.querySelector(".multiples_div .result").innerText = result;
  }).catch(err => {console.log(err)});
})

document.querySelector("#binary_submit").addEventListener('click', e => {
  var num = document.querySelector(".binary_div input").value;
  sendRequest("/binary", {"num" : num}).then(result => {
    document.querySelector(".binary_div .result").innerText = result;
  }).catch(err => {console.log(err)});
})

document.querySelector("#palindrom_submit").addEventListener('click', e => {

})

document.querySelector("#prime_submit").addEventListener('click', e => {

})

document.querySelector("#search_submit").addEventListener('click', e => {

})


function sendRequest(route, data){
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", route, true);
  xhttp.setRequestHeader("Content-Type", "application/json");
  xhttp.send(JSON.stringify(data));

  return new Promise((resolve, reject) => {
    xhttp.onreadystatechange = function() {
     if (this.readyState == 4 && this.status == 200) resolve(this.responseText);
     else if (this.status >= 400) reject(this.responseText);
   }
  });
}
