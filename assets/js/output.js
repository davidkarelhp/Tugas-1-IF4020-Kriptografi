let _5_letter_radio = document.getElementById("5-letter-radio");
let _5_letter_radio_label = document.querySelector("label[for='5-letter-radio']");
let no_space_radio = document.getElementById("no-space-radio");
let no_space_radio_label = document.querySelector("label[for='no-space-radio']");
let result_area = document.getElementById("result-area");

function chunkString(str, length) {
    return str.match(new RegExp('.{1,' + length + '}', 'g'));
}

_5_letter_radio.addEventListener("click", async (e) =>{
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

no_space_radio.addEventListener("click", async (e) =>{
    result_area.value = result_area.value.replace(/\s+/g, '');
    _5_letter_radio_label.classList.remove("disabled");
    no_space_radio_label.classList.add("disabled");
});