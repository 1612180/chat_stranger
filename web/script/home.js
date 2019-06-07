window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = sessionStorage.getItem('token');

    if (token !== null) {
        location.href = baseurl + '/web/chat_stranger/welcome_user'
    }

    let formSignIn = document.getElementById('formSignIn');
    formSignIn.addEventListener('submit', event => {
        event.preventDefault();

        let regname = document.getElementById('inputRegNameSignIn').value;
        let password = document.getElementById('inputPasswordSignIn').value;

        fetch(baseurl + '/api/chat_stranger/public/users/authenticate', {
            method: 'POST',
            body: JSON.stringify(
                {
                    regname: regname,
                    password: password
                }
            )
        })
            .then(res => res.json())
            .then(res => {
                console.log(res);
                if (res.code !== 206) {
                    return
                }
                sessionStorage.setItem('token', res.token);
                location.href = baseurl + '/web/chat_stranger/welcome_user'
            })
            .catch((err) => {
                console.log(err)
            })
    });

    let formSignUp = document.getElementById('formSignUp');
    formSignUp.addEventListener('submit', event => {
        event.preventDefault();

        let regname = document.getElementById('inputRegNameSignUp').value;
        let password = document.getElementById('inputPasswordSignUp').value;
        let fullname = document.getElementById('inputFullNameSignUp').value;

        fetch(baseurl + '/api/chat_stranger/public/users/register', {
            method: 'POST',
            body: JSON.stringify(
                {
                    regname: regname,
                    password: password,
                    fullname: fullname
                }
            )
        })
            .then(res => res.json())
            .then(res => {
                console.log(res);
                if (res.code !== 205) {
                    return
                }
                location.reload()
            })
    })
});
