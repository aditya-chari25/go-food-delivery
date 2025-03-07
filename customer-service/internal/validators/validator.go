package validator

import(
	"errors"
	"customer-service/internal/model"
	"fmt"
)

// validateOrders performs custom validation on the Orders struct
func ValidateOrder(order model.Orders) error {
    // Check if username is provided (seems to be a required field)
    if order.Username == "" {
        return errors.New("username is required")
    }
    if order.Userid <= 0 {
        return errors.New("valid user ID is required")
    }
    // Validate that we have at least one order item
    if len(order.Orders) == 0 {
        return errors.New("at least one order item is required")
    }
    // Validate each order item
    for i, item := range order.Orders {
        // Check restaurant information
        if item.RestaurantID == "" {
            return fmt.Errorf("order item #%d: restaurant ID is required", i+1)
        }
        
        if item.RestaurantName == "" {
            return fmt.Errorf("order item #%d: restaurant name is required", i+1)
        }
        
        // Check item name
        if item.Name == "" {
            return fmt.Errorf("order item #%d: item name is required", i+1)
        }
        
        // Check quantity
        if item.Quantity <= 0 {
            return fmt.Errorf("order item #%d: quantity must be greater than zero", i+1)
        }
        
        // Check price
        if item.Price <= 0 {
            return fmt.Errorf("order item #%d: price must be greater than zero", i+1)
        }
        
        // Check address
        if item.Address == "" {
            return fmt.Errorf("order item #%d: delivery address is required", i+1)
        }
    }
    // Calculate total order value (optional, for business logic validation)
    var totalAmount int
    for _, item := range order.Orders {
        totalAmount += item.Price * item.Quantity
    }
    // Optional: Minimum order value check
    if totalAmount < 100 { // Assuming minimum order value is 100 (adjust as needed)
        return errors.New("minimum order value is 100")
    }
    
    return nil
}