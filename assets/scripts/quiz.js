let quizData = [
    {
        question: "Does your site, app, and/or service collect information from people and use that information to display personalized ads?",
        a: "Yes",
        b: "No",
        correct: "a"
    },
    {
        question: "Does your site, app, and/or service collect information from people and send it to any third-party services (analytics, databases, files, etc.)?",
        a: "Yes",
        b: "No",
        correct: "a"
    },
    {
        question: "Does your site, app, and/or service collect information from people and use it to track people and show them your ads on other sites?",
        a: "Yes",
        b: "No",
        correct: "a"
    }
];

const quiz = document.getElementById('quiz');
const answers = document.querySelectorAll('.answer');
const question = document.getElementById('question');
const yesText = document.getElementById('yesText');
const noText = document.getElementById('noText');
const submitButton = document.getElementById('submit');

let currentQuiz = 0, counter = 0;

loadQuiz();

function loadQuiz() {
    deselectAnswers();

    const currentQuizData = quizData[currentQuiz];

    question.innerText = currentQuizData.question;
    yesText.innerText = currentQuizData.a
    noText.innerText = currentQuizData.b
}

function deselectAnswers() {
    answers.forEach(answerOption => answerOption.checked = false);
}

function getSelected() {
    let answer;

    answers.forEach(answerOption => {
        if (answerOption.checked) {
            answer = answerOption.id
        }
    })

    return answer
}

submitButton.addEventListener('click', () => {
    const selectedAnswer = getSelected();
    console.log(selectedAnswer);

    if (selectedAnswer) {
        if (selectedAnswer === quizData[currentQuiz].correct) {
            counter++;
        }

        currentQuiz++

        if (currentQuiz < quizData.length) {
            loadQuiz()
        } else if (counter === 1) {
            quiz.innerHTML = `
            <h2>You answered yes to ${counter} question</h2>
            <p>You need a cookie banner</p>
            <button onclick="location.reload()">Reload</button>
        `
        } else if (counter > 1) {
            quiz.innerHTML = `
            <h2>You answered yes to ${counter} questions</h2>
            <p>You need a cookie banner</p>
            <button onclick="location.reload()">Reload</button>
        `
        } else {
            quiz.innerHTML = `
            <h2>Congratulations! You don't need a cookie banner.</h2>
            <button onclick="location.reload()">Reload</button>
        `
        }
    }
})