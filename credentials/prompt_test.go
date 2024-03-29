package credentials_test

import (
	"testing"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/bool64/ctxd"
	"github.com/stretchr/testify/assert"
	"go.nhat.io/surveyexpect"

	"github.com/nhatthm/moneyloverprompt/credentials"
)

func TestCredentialsProvider_Username(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockSurvey     surveyexpect.Expector
		expectedResult string
	}{
		{
			scenario: "email is entered at the first time",
			mockSurvey: surveyexpect.Expect(func(s *surveyexpect.Survey) {
				s.ExpectPassword("Enter email (input is hidden) >").
					Answer("user@example.org")
			}),
			expectedResult: "user@example.org",
		},
		{
			scenario: "email is skipped at the first time and then entered",
			mockSurvey: surveyexpect.Expect(func(s *surveyexpect.Survey) {
				s.ExpectPassword("Enter email (input is hidden) >").Times(3)

				// Username is required, ask again.
				s.ExpectPassword("Enter email (input is hidden) >").
					Answer("user@example.org")
			}),
			expectedResult: "user@example.org",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			tc.mockSurvey(t).Start(func(stdio terminal.Stdio) {
				p := credentials.New(
					credentials.WithStdio(stdio),
					credentials.WithLogger(ctxd.NoOpLogger{}),
				)

				// 1st try: read from input.
				result := p.Username()

				assert.Equal(t, tc.expectedResult, result)

				// 2nd try: read from memory
				result = p.Username()

				assert.Equal(t, tc.expectedResult, result)
			})
		})
	}
}

func TestCredentialsProvider_UsernameInvalidInput(t *testing.T) {
	t.Parallel()

	s := surveyexpect.Expect(func(s *surveyexpect.Survey) {
		s.ExpectPassword("Enter email (input is hidden) >").
			Answer("\033X").Interrupted()
	})(t)

	expectedError := "error: could not read email {\"error\":{}}\n"

	s.Start(func(stdio terminal.Stdio) {
		l := &ctxd.LoggerMock{}
		p := credentials.New(
			credentials.WithStdio(stdio),
			credentials.WithLogger(l),
		)

		// 1st try: read from input.
		result := p.Username()

		assert.Empty(t, result)
		assert.Equal(t, expectedError, l.String())
	})
}

func TestCredentialsProvider_Password(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockSurvey     surveyexpect.Expector
		expectedResult string
	}{
		{
			scenario: "password is entered at the first time",
			mockSurvey: surveyexpect.Expect(func(s *surveyexpect.Survey) {
				s.ExpectPassword("Enter password (input is hidden) >").
					Answer("password")
			}),
			expectedResult: "password",
		},
		{
			scenario: "password is skipped at the first time and then entered",
			mockSurvey: surveyexpect.Expect(func(s *surveyexpect.Survey) {
				s.ExpectPassword("Enter password (input is hidden) >").Times(3)

				// Password is required, ask again.
				s.ExpectPassword("Enter password (input is hidden) >").
					Answer("password")
			}),
			expectedResult: "password",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			tc.mockSurvey(t).Start(func(stdio terminal.Stdio) {
				p := credentials.New(
					credentials.WithStdio(stdio),
					credentials.WithLogger(ctxd.NoOpLogger{}),
				)

				// 1st try: read from input.
				result := p.Password()

				assert.Equal(t, tc.expectedResult, result)

				// 2nd try: read from memory
				result = p.Password()

				assert.Equal(t, tc.expectedResult, result)
			})
		})
	}
}

func TestCredentialsProvider_PasswordInvalidInput(t *testing.T) {
	t.Parallel()

	s := surveyexpect.Expect(func(s *surveyexpect.Survey) {
		s.ExpectPassword("Enter password (input is hidden) >").
			Answer("\033X").Interrupted()
	})(t)

	expectedError := "error: could not read password {\"error\":{}}\n"

	s.Start(func(stdio terminal.Stdio) {
		l := &ctxd.LoggerMock{}
		p := credentials.New(
			credentials.WithStdio(stdio),
			credentials.WithLogger(l),
		)

		// 1st try: read from input.
		result := p.Password()

		assert.Empty(t, result)
		assert.Equal(t, expectedError, l.String())
	})
}
