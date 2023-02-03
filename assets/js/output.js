let _5_letter_radio = document.getElementById("5-letter-radio");
let _5_letter_radio_label = document.querySelector("label[for='5-letter-radio']");
let no_space_radio = document.getElementById("no-space-radio");
let no_space_radio_label = document.querySelector("label[for='no-space-radio']");
let result_area = document.getElementById("result-area");
let download_output = document.getElementById("download-output");

function chunkString(str, length) {
    return str.match(new RegExp('.{1,' + length + '}', 'g'));
}

_5_letter_radio.addEventListener("click", (e) =>{
    chunks = chunkString(result_area.value, 5);
    let newText = "";
    let l = chunks?.length ?? 0;
    chunks?.forEach((element, idx) => {
        newText += element;
        if (idx < l - 1) {
            newText += " "
        }
    });
    result_area.value = newText;
    _5_letter_radio_label.classList.add("disabled");
    no_space_radio_label.classList.remove("disabled");
});

no_space_radio.addEventListener("click", (e) =>{
    result_area.value = result_area.value.replace(/\s+/g, '');
    _5_letter_radio_label.classList.remove("disabled");
    no_space_radio_label.classList.add("disabled");
});

download_output.addEventListener("click", (e) =>{
    if (result_area.value == null || result_area.value == "") {
        alert("Output Text is empty");
    } else {
        download("output.txt", result_area.value);
    }
    
});

function download(filename, text) {
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);
  
    element.style.display = 'none';
    document.body.appendChild(element);
  
    element.click();
  
    document.body.removeChild(element);
}