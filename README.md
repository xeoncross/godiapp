## Go Dependency injection App

Test application design + folder package structure to achieve a clean web
application that is testable and without strong uncoupling.

- Version 2: https://www.reddit.com/r/golang/comments/91a54y/proper_structure_for_a_dependencyinjection_based/
- Version 1: https://www.reddit.com/r/golang/comments/8zn2ti/how_to_handle_dependency_injection_and_mocking/ (old)

## Questions:

### Testing + Mocks

Currently the only test (`/mock`) in here is pointlessly testing itself. How should this be refactored? Should the tests be placed next to the implementations and only *use* the `/mocks`?

### Scale

How will this approach work with 20+ domain objects (users,likes,comments,etc...)? At the moment the `/mysql` implementation of `UserService` has all the possible database queries defined on itself with no separation. Should `UserService` become a monolithic `StoreService` or should we add 20+ more `___Service` interfaces?

This same question goes for the `http` handlers and other possible API's such as GraphQL.

#### Example: Mattermost (OS Slack Alternative)

Mattermost is an example that defines a base [store](https://github.com/mattermost/mattermost-server/blob/master/store/store.go) that wraps all the individual entity store interfaces. Then provides [sql and other](https://github.com/mattermost/mattermost-server/blob/master/store/sqlstore/store.go) concrete implementations of these interfaces for actual storage. This includes a [mock](https://github.com/mattermost/mattermost-server/blob/master/store/storetest/store.go) store. They have 20+ entities.

## Design choices and Articles:

- [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
- [Interface Location](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)
- [Wrap sql.DB and sql.Tx](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)
- [HTTP Closures](https://gist.github.com/tsenart/5fc18c659814c078378d)
- [Simple Go Server](https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7)
- [Dependency injection](https://www.alexedwards.net/blog/organising-database-access#using-an-interface)
- [HTTP middleware](https://gist.github.com/Xeoncross/372bb42c24b1cb37664c377d018dd5cb)
- [Discussion on passing db to http.Handlers](https://www.reddit.com/r/golang/comments/5vsz2t/what_is_the_best_way_to_pass_a_db_to_web_handlers/)
- [Should I mock the DB?](https://www.reddit.com/r/golang/comments/6n3m4w/is_there_a_good_use_case_for_mocking_a_db/)
- [http.Handler wrapping](https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702)
