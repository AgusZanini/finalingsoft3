Feature('crud');

Scenario('test crud operations',  ({ I }) => {
    I.amOnPage("https://frontend-fuz7myfmqq-uc.a.run.app")
    I.fillField('#num1', '5');
    I.fillField('#num2', '10');
    I.click('#multiply');
    I.see('Resultado: 50', 'h3');
    I.fillField('#idinsert', '7')
    I.see('Operación: {"Id":7,"number1":4,"number2":3,"result":12}', 'h3')
});
