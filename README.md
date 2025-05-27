# Go-eat with Monolithic Architecture

Welcome to Go Eat, my fictional restaurant. Please, take a seat; our journey is about to begin.

This space chronicles the evolution from a small app to a scaled solution, offering a practical template of ideas born from my commercial ventures. I'm doing this primarily as a safe environment to prototype new features without affecting my primary application. It's a place to explore concepts, get feedback, and ideally, create something beneficial for the wider community.

## Operations

Efficiency is a core principle at Go Eat. We've found that customers who appreciate streamlined, automated processes truly embrace our model. By serving a fixed menu daily with no deviations, we enable our staff to concentrate fully on their primary roles, freeing them from tedious and repetitive administrative tasks.

### Booking

We handle customer bookings through a third-party app. Upon a successful reservation, our webhook endpoint immediately notifies the kitchen.

### Kitchen

We handle customer bookings through a third-party app. Upon a successful reservation, our webhook endpoint immediately notifies the kitchen. Each customer reservation automatically schedules the corresponding orders for the kitchen. Our automated kitchen stock system then proactively manages inventory, triggering outbound requests to suppliers as needed. When plates are ready, a notification is seamlessly dispatched to the service staff.

## The House Style

Why Go?
This entire system is built with Go, a language I deeply appreciate. It offers ample type safety without being overly rigid and consistently prioritizes simplicity. While every language has its trade-offs, Go's pragmatic approach aligns perfectly with my goals here.

You'll notice that while I generally adhere to Go's idiomatic patterns, sometimes I diverge. This is overwhelmingly for pragmatic reasons, as I'm not just building simple packages. Stepping outside the strict idiomatic Go layout occasionally serves the greater interest of overall simplicity for this project.

## Designed for Future Growth

Every architectural decision is made with an eye on future scalability. I'm not building for a million-plus customers from day one, but I am consciously avoiding code that would complicate scaling down the line.

## Monolith

Every architectural decision is made with an eye on future scalability. I'm not building for a million-plus customers from day one, but I am consciously avoiding code that would complicate scaling down the line.

A critical juncture in scaling any business or application arises when you hire enough developers to allow for specialized effort. While coordinating one to five developers on a shared codebase is typically straightforward, exceeding that number often makes a single, large codebase unwieldy without individual developer specialization. At this point, breaking code into distinct repositories seems like the obvious choice, necessitating either a complete refactor or a smooth transition to a service-oriented architecture.

Go's straightforward package system is inherently well-suited for separating concerns across distinct repositories. However, the practical reality of continuously updating local copies, coupled with the friction of deploying and versioning multiple services, can become a significant drag on development velocity.

This is why, for now, we embrace the monolith. A monolith offers inherent simplicity, especially with Go's package structure, which can elegantly support a service-like architecture within a single codebase. It also makes deploying (and reverting) a single process exceptionally easy.

Crucially, we will build and lay out this monolith in a way that makes extracting components straightforward in the future. For instance, if our kitchen component expands to include take-away services, we can refactor just that part into an external repository and dedicate a specialized team to it. However, slowing down development today by externalizing it prematurely makes no sense.

Our choice of a monolith stems from several key considerations:

This application is intended to grow into a much larger system.
Development begins with a small team (1-4 developers) who can easily collaborate.
We have a low tolerance for administrative overhead that multiple independent services often introduce.

### Package Layout

The initial package layout of this project deliberately reflects the various domains and concerns within our application and business operations. Each of these packages could, in time, become an independent application or service.

In fact, as the business scales, the intention is to spin out these packages as standalone services. For example, we might eventually hire a dedicated development team to work exclusively on the menu system. This thoughtful internal separation will significantly ease that transition if and when it occurs, while also maintaining Go-like separation of functionality in the interim.

```
./goeat/booking
./goeat/menu
./goeat/staff
./goeat/service
./goeat/kitchen
./goeat/README.md
./goeat/main.go
./goeat/go.mod
```

### Domain Boundaries

A common pitfall that plagues companies as they scale is what I call "monolith mess"—code execution paths become so deeply intertwined that the only pragmatic path to further scaling is a costly, full application rewrite. While a rewrite at least clarifies the desired functionality, it's an inefficient way to grow.

To proactively avoid this, I'm leveraging Go's internal package qualifier. For instance, when developing functions within the kitchen domain, I want to strictly prevent direct calls into the staff domain, even if it seems convenient at the moment.

Go's default package export mechanism, based on capitalized type names, is often too permissive at the service or domain level for my specific needs. I require a stronger enforcement to protect inter-service package exports, while still allowing full usage within a service.

To achieve this, I'm adopting a clear convention: a dedicated publicapi.go file at the root of each major (service-like) package tree. By explicitly defining all external-facing functions in this publicapi.go file, and enclosing all other internal functions under the internal qualifier, I create a robust system of enforced boundaries and clear APIs. This ensures that even within the monolith, our domains remain decoupled and ready for independent extraction.

```
./goeat/kitchen/
./goeat/kitchen/internal/rota.go

./goeat/staff/
./goeat/staff/publicapi.go
./goeat/staff/internal/calendar/main.go
```

My `kitchen` service needs to update its `rota` from the `calendar` package in the `staff` service. So I create a public API route into `staff` like this:

```
// ./goeat/staff/publicapi.go

package staff

import "goeat/staff/internal/calendar"

var (
        GetKitchenRota = func() []string {
                return calendar.GetRota("kitchen")
        }
)
```

It's really just a local version of an http API. Which will make a future transition easier to do.

And calling this from the `kitchen` service looks like this:

``` 
// ./goeat/kitchen/internal/rota.go

package rota

import (
	"goeat/staff"
)

func fetchRota() []string {

	rota := staff.GetKitchenRota()

	// do something with the rota

	return rota
}
```

A readability bonus is the name of the import being `staff`. The domain is much clearer than using a sub-package export, which would result in `calendar.GetRota`.


### Testing

A particularly effective benefit of implementing these public "API" methods as Go vars is the ability to isolate testing between services. It becomes trivial to mock the responses from, say, the staff service, entirely independently of that package's implementation. I've found this to be the simplest approach for emulating distinct services within a monolith, avoiding the code gymnastics often associated with interfaces or channels in other mocking strategies.

For example, a simple test within the kitchen service might look like this:

``` 
//go:build test

package rota

import (
	"goeat/staff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRota(t *testing.T) {

	mockRota := []string{"A", "B"}
	mockStaffGetKitchenRota := func() []string {
		return mockRota
	}

	// override the 'api' response with our mock function
	staff.GetKitchenRota = mockStaffGetKitchenRota

	assert.Equal(t, fetchRota(), mockRota, "expect A, B")
}
```

### Summary

My decision to develop this application as a monolith is driven by a set of specific, pragmatic use-case factors.

Go's inherent simplicity regarding packages makes it remarkably easy to establish clear domain boundaries within the codebase. This clarity, in turn, makes the entire system simpler to conceptualize and manage.

Simplicity is paramount, especially when working with a small team, or even as a solo developer; complexity can quickly erode motivation. In my experience, development speed is often less about raw CPU cycles and more about establishing efficient processes.

At this initial scale, a single, deployable binary is inherently easier to manage and test. While Go is indeed well-suited for a discrete service architecture, the associated overhead—managing multiple repositories, complex continuous integration pipelines, and intricate infrastructure—is simply too much for our current stage. This pragmatic approach allows us to focus on rapid feature development without being bogged down by unnecessary architectural complexity.

