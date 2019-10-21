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
