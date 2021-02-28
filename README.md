Task for Backend Developer:
Develop the API for an online e-commerce store selling fruit, which contains the following features:
1.  User sign-up and login.
2.  Browse the following products:
    -  Apples
    -  Bananas
    -  Pears
    -  Oranges
3.  Manage coupon codes through an admin.
4.  Go to checkout.
5.  Mocked purchase (a payment gateway is not required, but a route must exist in the backend validating the payment).
6.  An address does not need to be entered.
 
Checkout Rules
    1.  If 7 or more apples are added to the cart, a 10% discount is applied to all apples.
    2.  For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set.
    3.  These sets must be added to their own cart item entry.
    4.  If pears or bananas already exist in the cart, this discount must be recalculated when new pears or bananas are added.
    5.  A coupon code can be used to get a 30% discount on oranges, if applied to the cart, otherwise oranges are full price.
    6.  Can only be applied once.
    7.  Has a configurable expiry timeout (10 seconds default) once generated.

Requirements
1.  Architecture diagrams.
2.  Backend RESTful web service written in Go.
3.  Rudimentary client/fontend to demonstrate and test the API.
4.  Use up-to-date technology.
