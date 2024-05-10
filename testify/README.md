# testify

`testify` is most commonly used unit testing framework for GoLang.

`assert` & `require` are packages

- key difference b/t assert & require

`require` : if an assertion fails, it immediately stops the execution of the test function by calling `t.Fatal` or `t.FailNow`. 

Used for preventing further execution of potentially faulty code.

assertion functions provided by `assert` :

1) `assert.Equal(t, expected, actual)` : checks if 2 values are Equal
2) `assert.NotEqual(t, expected, actual)` : checks if 2 values are Not Equal
3) `assert.True(t, condition)` : checks if condition/value is true
4) `assert.False(t, condition)` : checks if condition/value is false
5) `assert.Nil(t, value)` :  checks if a value is nil
6) `assert.NotNil(t, value)` :  checks if a value is not nil
7) `assert.Error(t, err)` : checks if a function call returns an error
8) `assert.NoError(t, err)` : checks if a function call does not return an error

assertion functions provided by `require` :

1) require.Equal(t, expected, actual)
2) require.NotEqual(t, expected, actual)
3) require.True(t, condition)
4) require.False(t, condition)
5) require.Nil(t, value)
6) require.NotNil(t, value)
7) require.Error(t, err)
8) require.NoError(t, err)




