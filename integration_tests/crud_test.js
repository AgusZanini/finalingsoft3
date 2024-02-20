Feature('crud');

Scenario('test crud operations',  ({ I }) => {
    I.amOnPage("https://frontend-fuz7myfmqq-uc.a.run.app")
    I.fillField('input[type=number]:nth-child(1)', '5');
    I.fillField('input[type=number]:nth-child(2)', '10');
    I.click('button');
    I.see('Resultado: 50', 'h3');
});
