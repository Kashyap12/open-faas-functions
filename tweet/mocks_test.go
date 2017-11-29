// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package function

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"sync"
)

var (
	lockTwitterPosterMockPostTweet sync.RWMutex
)

// TwitterPosterMock is a mock implementation of TwitterPoster.
//
//     func TestSomethingThatUsesTwitterPoster(t *testing.T) {
//
//         // make and configure a mocked TwitterPoster
//         mockedTwitterPoster := &TwitterPosterMock{
//             PostTweetFunc: func(status string, v url.Values) (anaconda.Tweet, error) {
// 	               panic("TODO: mock out the PostTweet method")
//             },
//         }
//
//         // TODO: use mockedTwitterPoster in code that requires TwitterPoster
//         //       and then make assertions.
//
//     }
type TwitterPosterMock struct {
	// PostTweetFunc mocks the PostTweet method.
	PostTweetFunc func(status string, v url.Values) (anaconda.Tweet, error)

	// calls tracks calls to the methods.
	calls struct {
		// PostTweet holds details about calls to the PostTweet method.
		PostTweet []struct {
			// Status is the status argument value.
			Status string
			// V is the v argument value.
			V url.Values
		}
	}
}

// PostTweet calls PostTweetFunc.
func (mock *TwitterPosterMock) PostTweet(status string, v url.Values) (anaconda.Tweet, error) {
	if mock.PostTweetFunc == nil {
		panic("moq: TwitterPosterMock.PostTweetFunc is nil but TwitterPoster.PostTweet was just called")
	}
	callInfo := struct {
		Status string
		V      url.Values
	}{
		Status: status,
		V:      v,
	}
	lockTwitterPosterMockPostTweet.Lock()
	mock.calls.PostTweet = append(mock.calls.PostTweet, callInfo)
	lockTwitterPosterMockPostTweet.Unlock()
	return mock.PostTweetFunc(status, v)
}

// PostTweetCalls gets all the calls that were made to PostTweet.
// Check the length with:
//     len(mockedTwitterPoster.PostTweetCalls())
func (mock *TwitterPosterMock) PostTweetCalls() []struct {
	Status string
	V      url.Values
} {
	var calls []struct {
		Status string
		V      url.Values
	}
	lockTwitterPosterMockPostTweet.RLock()
	calls = mock.calls.PostTweet
	lockTwitterPosterMockPostTweet.RUnlock()
	return calls
}