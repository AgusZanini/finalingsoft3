Feature('crud');

Scenario('Multiply two numbers and verify result', ({ I }) => {
    I.amOnPage('/');
    I.fillField('#num1', '5');
    I.fillField('#num2', '3');
    I.click('#multiply');
    I.wait(5);
    I.see('Resultado: 15');
  });

Scenario('Get an operation by id', ({ I }) => {
    I.amOnPage("/")
    I.fillField('#idinsert', '7');
    I.click('#getbyid');
    I.see('Operación: {"Id":7,"number1":4,"number2":3,"result":12}');
})
